package tools

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/model"
	"encoding/json"
	"io"
	"io/ioutil"
)

// VideoToRsp 传入视频结构体和作者结构体
func VideoToRsp(v *model.Video, a *model.User) *api.VideoDetailRespond {
	var rsp api.VideoDetailRespond
	rsp.FavoriteCount = int64(v.FavoriteCount)
	rsp.PlayURL = v.PlayURL
	rsp.CommentCount = int64(v.CommentCount)
	rsp.ID = int64(v.VideoID)
	rsp.CoverURL = v.CoverURL
	rsp.Title = v.Title

	rsp.Author.BackgroundImage = a.BackgroundImage
	rsp.Author.ID = int64(a.UserID)
	rsp.Author.Avatar = a.Avatar
	rsp.Author.Name = a.Name
	rsp.Author.FavoriteCount = a.FavoriteCount
	rsp.Author.FollowCount = a.FollowCount
	rsp.Author.FollowerCount = a.FollowerCount
	rsp.Author.Signature = a.Signature
	rsp.Author.WorkCount = a.WorkCount
	rsp.Author.TotalFavorited = a.TotalFavorite
	return &rsp
}

func UserToRsp(a *model.User) *api.UserDetailRespond {
	var rsp api.UserDetailRespond
	rsp.BackgroundImage = a.BackgroundImage
	rsp.ID = int64(a.UserID)
	rsp.Avatar = a.Avatar
	rsp.Name = a.Name
	rsp.FavoriteCount = a.FavoriteCount
	rsp.FollowCount = a.FollowCount
	rsp.FollowerCount = a.FollowerCount
	rsp.Signature = a.Signature
	rsp.WorkCount = a.WorkCount
	rsp.TotalFavorited = a.TotalFavorite
	return &rsp
}

func CommentToRsp(c *model.Comment) *api.CommentDetailRsp {
	var r api.CommentDetailRsp
	r.ID = int64(c.CommentID)
	r.CreateDate = c.CreateDate
	r.Content = c.Content
	return &r
}

func CommentsToRsp(c []model.Comment) []api.CommentDetailRsp {
	r := make([]api.CommentDetailRsp, len(c))
	for i, v := range c {
		r[i] = *CommentToRsp(&v)
	}
	return r
}

func MessagestoRsp(m []model.ChatMessage) []api.MessageDetailRsp {
	r := make([]api.MessageDetailRsp, len(m))
	for i, v := range m {
		r[i] = api.MessageDetailRsp{
			FromUserID: int64(v.FromUserID),
			ToUserID:   int64(v.ToUserID),
			ID:         int64(v.ChatMessageID),
			Content:    v.Content,
			CreateTime: v.CreateTime,
			//CreateTime: time.Unix(v.CreateTime, 0).Format(time.DateTime),
		}
	}
	return r
}

func FileToRsp(body io.ReadCloser) (*api.FileUploadsRsp, error) {
	//从body中读出数据
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	//反序列化到rsp结构体中
	var rsp api.FileUploadsRsp
	err = json.Unmarshal(data, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, nil
}
