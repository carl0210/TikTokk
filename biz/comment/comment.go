package comment

import (
	"TikTokk/api"
	"TikTokk/model"
	"TikTokk/store"
	"TikTokk/utils"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CommentBiz interface {
	Delete(ctx context.Context, commentID, videoID uint, username string) error
	Create(ctx context.Context, videoID uint, username, commentText string) (*api.CommentActionRsp, error)
	List(ctx context.Context, videoID uint) (*api.CommentListRsp, error)
}

type BComment struct {
	ds store.DataStore
}

var _ CommentBiz = (*BComment)(nil)

func New(s store.DataStore) *BComment {
	return &BComment{ds: s}
}

func (b BComment) List(ctx context.Context, videoID uint) (*api.CommentListRsp, error) {
	var rsp api.CommentListRsp
	//store 获取所有评论,按倒序
	list, err := b.ds.Comment().List(ctx, videoID)
	if err != nil {
		return &rsp, err
	}
	//返回
	rList := make([]api.CommentDetailRsp, len(list))
	for i, v := range list {
		rList[i] = *utils.CommentToRsp(&v)
		u, err := b.ds.Users().GetByID(ctx, v.UserId)
		if err != nil {
			return &rsp, err
		}
		rList[i].User = *utils.UserToRsp(u)
	}
	rsp.CommentList = rList
	return &rsp, nil
}

func (b BComment) Delete(ctx context.Context, commentID, videoID uint, username string) error {
	//得到comment具体信息,对比申请删除是否同作者相同
	comment, err := b.ds.Comment().GetByCommentID(ctx, commentID)
	if err != nil {
		return err
	}
	if comment.UserName != username {
		return fmt.Errorf("非评论发布者")
	}
	//获取视频信息得到评论数
	v, err := b.ds.Videos().GetByVideoID(ctx, videoID)
	if err != nil {
		return err
	}
	//创建事务,删除记录,视频评论数-1
	f := func(tx *gorm.DB) error {
		if err := tx.Table("comments").Delete(&model.Comment{Id: commentID}).Error; err != nil {
			return err
		}
		if err := tx.Table("videos").Model(&model.Video{VideoId: videoID}).Update(
			"comment_count", v.CommentCount-1).Error; err != nil {
			return err
		}
		return nil
	}
	if err := b.ds.Comment().Transaction(ctx, f); err != nil {
		return err
	}
	return nil
}

func (b BComment) Create(ctx context.Context, videoID uint, username, commentText string) (*api.CommentActionRsp, error) {
	var rsp api.CommentActionRsp
	//得到视频总评论数
	v, err := b.ds.Videos().GetByVideoID(ctx, videoID)
	if err != nil {
		return &rsp, err
	}
	//得到作者信息,并转化
	u, err := b.ds.Users().GetByName(ctx, username)
	if err != nil {
		return &rsp, err
	}
	rU := utils.UserToRsp(u)
	//创建者同登录账号相同,故未关注
	rU.IsFollow = false
	//创建时间为服务器当前时间
	nowTime := time.Now().Format("01-02")
	//创建结构体
	c := model.Comment{VideoId: videoID, CreateDate: nowTime, Content: commentText, UserName: username, UserId: u.UserId}
	//创建事务,创建记录,视频评论数+1
	f := func(tx *gorm.DB) error {
		if err := tx.Table("comments").Create(&c).Error; err != nil {
			return err
		}
		if err := tx.Table("videos").Model(&model.Video{VideoId: videoID}).Update(
			"comment_count", v.CommentCount+1).Error; err != nil {
			return err
		}
		return nil
	}
	if err := b.ds.Comment().Transaction(ctx, f); err != nil {
		return &rsp, err
	}
	//得到创建完成的信息
	com, err := b.ds.Comment().GetByUName(ctx, videoID, username, nowTime)
	if err != nil {
		return &rsp, err
	}
	rsp.Comment = *utils.CommentToRsp(com)
	rsp.Comment.User = *rU
	return &rsp, err

}
