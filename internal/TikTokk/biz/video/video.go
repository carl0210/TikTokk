package video

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/err"
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/TikTokk/var"
	tencent_cos "TikTokk/internal/pkg/tencent-cos"
	"TikTokk/tools"
	t_ffmpeg "TikTokk/tools/t-ffmpeg"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoBiz interface {
	GetVideoFeedList(ctx context.Context, userID uint, latestTime int64) (rsp *api.VideoFeedListRsp, err error)
	PublishAction(ctx *gin.Context, file *multipart.FileHeader, title string, userID uint) error
	PublishList(ctx context.Context, userID uint) (*api.VideoPublishListRsp, error)
}

type BVideo struct {
	ds store.DataStore
}

var _ VideoBiz = (*BVideo)(nil)

func New(s store.DataStore) *BVideo {
	return &BVideo{ds: s}
}

func (b *BVideo) GetVideoFeedList(ctx context.Context, userID uint, latestTime int64) (rsp *api.VideoFeedListRsp, err error) {
	//获取限制返回视频的最新投稿时间
	//获取视频
	list, err := b.ds.Videos().Feed(ctx, FeedLen, time.Unix(latestTime, 0))
	if err != nil {
		return &api.VideoFeedListRsp{}, err
	}
	//根据获取的视频列表构建返回的视频列表
	NextTime := latestTime
	videoList := make([]api.VideoDetailRespond, len(list))
	//得到登录用户的信息，以获得对各视频的喜欢状态和其作者的关注状态
	var u *model.User
	if userID != 0 {
		//如果存在登录用户
		u, err = b.ds.Users().Get(ctx, &model.User{UserID: uint(userID)})
		if err != nil {
			return &api.VideoFeedListRsp{}, err
		}
	} else {
		//未登录则将名字置为“”表示游客
		u = &model.User{Name: ""}
	}

	for i, v := range list {
		//得到作者信息
		author, err := b.ds.Users().Get(ctx, &model.User{UserID: v.AuthorId})
		if err != nil {
			return &api.VideoFeedListRsp{}, err
		}
		//构建返回视频列表
		videoList[i] = *tools.VideoToRsp(&v, author)
		//playUrl, err := minio.GetObject(ctx, "dev", v.PlayKey, 1*time.Hour)
		//if err != nil {
		//	return &api.VideoFeedListRsp{}, err
		//}
		videoList[i].PlayURL = v.PlayKey
		log.Println(v.PlayKey)
		//得到视频的喜欢状态
		//如果为游客，则默认未喜欢
		if len(u.Name) == 0 {
			videoList[i].IsFavorite = false
		} else {
			//若登录账号则拉取关系
			relFavorite, err := b.ds.VideoFavoriteRelation().FirstOrCreate(ctx, v.VideoID, u.UserID, u.Name)
			if err != nil {
				return &api.VideoFeedListRsp{}, err
			}
			videoList[i].IsFavorite = relFavorite.ISFavorite
		}
		//得到其作者的关注状态
		//获取用户对作者关注关系
		if len(u.Name) == 0 {
			//如果为游客,则置为未关注
			videoList[i].Author.IsFollow = false
		} else {
			relFollow, err := b.ds.UserFollowRelation().FirstOrCreate(ctx, u.UserID, author.UserID, u.Name, author.Name)
			if err != nil {
				return &api.VideoFeedListRsp{}, err
			}
			videoList[i].Author.IsFollow = relFollow.IsFollow
		}
		//获取该切片视频中最早时间
		if v.UpdatedAt.Unix() < NextTime {
			NextTime = v.UpdatedAt.Unix()
		}

	}
	return &api.VideoFeedListRsp{VideoList: videoList, NextTime: NextTime}, nil
}

func (b *BVideo) PublishAction(ctx *gin.Context, file *multipart.FileHeader, title string, userID uint) error {
	// 判断文件大小是否超过限制
	if file.Size > _var.VIDEO_MAX_LENGTH_LIMIT {
		return err.EXCEED_MAXiMUN_VIDEO_LENGTH_ERROR
	}
	//构建文件名: 用户名-上传时间戳-文件名,创建路径
	fileNameLatest := fmt.Sprintf("%d-%d-%s", userID, time.Now().Unix(), filepath.Base(file.Filename))
	// err = minio.PutObject(ctx, "dev", fileNameLatest, reader, file.Size)
	// 保存为临时文件
	fd, err := file.Open()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer fd.Close()
	temp, err := os.CreateTemp("./tmp/", "tmp.*.mp4")
	if err != nil {
		log.Println("临时文件创建失败！, err = ", err)
		return err
	}
	defer os.Remove(temp.Name())
	defer temp.Close()
	// 压缩视频
	err = t_ffmpeg.CompressVideo(fd, temp.Name())
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// 上传视频
	videoPlayUrl, err := tencent_cos.UploadVideo(ctx, fileNameLatest, temp)
	log.Println("VideoPlayUrl = ", videoPlayUrl, err)
	if err != nil {
		return err
	}

	// 生成缩略图
	photoFileName := fmt.Sprintf("%d-%d.jpg", userID, time.Now().Unix())
	photoUrl, err := tencent_cos.UploadVideoCoverByHeader(ctx, photoFileName, temp)
	if err != nil {
		log.Println(err)
		return err
	}
	//更新用户视频数和创建视频记录
	//获取作者信息
	u, err := b.ds.Users().Get(ctx, &model.User{UserID: userID})
	if err != nil {
		return err
	}
	//创建视频记录
	v := model.Video{
		//PlayKey:  fileNameLatest,
		PlayKey:  videoPlayUrl,
		Title:    title,
		AuthorId: u.UserID,
		CoverURL: photoUrl,
	}
	err = b.ds.Videos().Create(ctx, &v)
	if err != nil {
		return err
	}
	//作者作品数+1
	newU := u
	newU.WorkCount += 1
	err = b.ds.Users().Update(ctx, u.Name, newU)
	if err != nil {
		return err
	}
	return nil

}

func (b *BVideo) PublishList(ctx context.Context, userID uint) (*api.VideoPublishListRsp, error) {
	var rsp api.VideoPublishListRsp
	//得到userID的User结构体
	u, err := b.ds.Users().Get(ctx, &model.User{UserID: userID})
	if err != nil {
		return &rsp, err
	}
	//根据userID查找所有的视频
	list, err := b.ds.Videos().ListAllVideoByAuthorIDLen(ctx, userID, u.WorkCount)
	if err != nil {
		return &rsp, err
	}
	//构建[]VideoDetailRsp
	rspList := make([]api.VideoDetailRespond, len(list))
	for i, v := range list {
		rspList[i] = *tools.VideoToRsp(&v, u)
		//得到点赞关系
		rel, err := b.ds.VideoFavoriteRelation().FirstOrCreate(ctx, v.VideoID, u.UserID, u.Name)
		if err != nil {
			return &rsp, err
		}
		rspList[i].IsFavorite = rel.ISFavorite
		//作者就是用户自身,则设置为未关注
		rspList[i].Author.IsFollow = false
	}
	rsp.VideoList = rspList
	return &rsp, nil
}
