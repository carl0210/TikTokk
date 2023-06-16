package biz

import (
	"TikTokk/biz/comment"
	"TikTokk/biz/message"
	"TikTokk/biz/relation/favorite"
	"TikTokk/biz/relation/follow"
	"TikTokk/biz/user"
	"TikTokk/biz/video"
	"TikTokk/store"
)

type IBiz interface {
	Users() user.UserBiz
	Videos() video.VideoBiz
	FavoriteRel() favorite.FavoriteRelationBiz
	Comment() comment.CommentBiz
	Follow() follow.FollowRelationBiz
	Message() message.MessageBiz
}

type biz struct {
	ds store.DataStore
}

var _ IBiz = (*biz)(nil)

func NewBiz(db store.DataStore) *biz {
	return &biz{ds: db}
}

func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}

func (b *biz) Videos() video.VideoBiz {
	return video.New(b.ds)
}

func (b *biz) FavoriteRel() favorite.FavoriteRelationBiz {
	return favorite.New(b.ds)
}

func (b *biz) Comment() comment.CommentBiz {
	return comment.New(b.ds)
}

func (b *biz) Follow() follow.FollowRelationBiz {
	return follow.New(b.ds)
}

func (b *biz) Message() message.MessageBiz {
	return message.New(b.ds)
}
