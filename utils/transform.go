package utils

import (
	"TikTokk/api"
	"TikTokk/model"
)

// VideoToRsp 传入视频结构体和作者结构体
func VideoToRsp(v *model.Video, a *model.User) *api.VideoDetailRespond {
	var rsp api.VideoDetailRespond
	rsp.FavoriteCount = int64(v.FavoriteCount)
	rsp.PlayURL = v.PlayURL
	rsp.CommentCount = int64(v.CommentCount)
	rsp.ID = int64(v.VideoId)
	rsp.CoverURL = v.CoverURL
	rsp.Title = v.Title

	rsp.Author.BackgroundImage = a.BackgroundImage
	rsp.Author.ID = int64(a.UserId)
	rsp.Author.Avatar = a.Avatar
	rsp.Author.Name = a.Name
	rsp.Author.FavoriteCount = int64(a.FavoriteCount)
	rsp.Author.FollowCount = int64(a.FollowCount)
	rsp.Author.FollowerCount = int64(a.FollowerCount)
	rsp.Author.Signature = a.Signature
	rsp.Author.WorkCount = int64(a.WorkCount)
	rsp.Author.TotalFavorited = a.TotalFavorited
	return &rsp
}

func UserToRsp(a *model.User) *api.UserDetailRespond {
	var rsp api.UserDetailRespond
	rsp.BackgroundImage = a.BackgroundImage
	rsp.ID = int64(a.UserId)
	rsp.Avatar = a.Avatar
	rsp.Name = a.Name
	rsp.FavoriteCount = int64(a.FavoriteCount)
	rsp.FollowCount = int64(a.FollowCount)
	rsp.FollowerCount = int64(a.FollowerCount)
	rsp.Signature = a.Signature
	rsp.WorkCount = int64(a.WorkCount)
	rsp.TotalFavorited = a.TotalFavorited
	return &rsp
}

func CommentToRsp(c *model.Comment) *api.CommentDetailRsp {
	var r api.CommentDetailRsp
	r.ID = int64(c.Id)
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

func MessagestoRsp(m []model.Chat_Message) []api.MessageDetailRsp {
	r := make([]api.MessageDetailRsp, len(m))
	for i, v := range m {
		r[i] = api.MessageDetailRsp{
			FromUserID: int64(v.FromUserID),
			ToUserID:   int64(v.ToUserID),
			ID:         int64(v.Id),
			Content:    v.Content,
			CreateTime: v.CreateTime,
			//CreateTime: time.Unix(v.CreateTime, 0).Format(time.DateTime),
		}
	}
	return r
}
