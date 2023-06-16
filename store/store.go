package store

import (
	"TikTokk/store/relation"
	"gorm.io/gorm"
	"sync"
)

type DataStore interface {
	Users() UserStore
	Videos() VideoStore
	UserFollowRelation() relation.IUserFollowRelation
	VideoFavoriteRelation() relation.IVideoFavoriteRelation
	Comment() CommentStore
	Message() MessageStore
}

type SData struct {
	db *gorm.DB
}

var (
	once sync.Once
	S    *SData
)

var _ DataStore = (*SData)(nil)

func NewStore(db *gorm.DB) *SData {
	once.Do(func() {
		S = &SData{db: db}
	})
	return S
}

func (s *SData) Videos() VideoStore {
	return NewVideos(s.db)
}

func (s *SData) Users() UserStore {
	return NewUsers(s.db)
}

func (s *SData) UserFollowRelation() relation.IUserFollowRelation {
	return relation.NewSUserFollowRelation(s.db)
}

func (s *SData) VideoFavoriteRelation() relation.IVideoFavoriteRelation {
	return relation.NewSVideoFavoriteRelation(s.db)
}

func (s *SData) Comment() CommentStore {
	return NewSComment(s.db)
}

func (s *SData) Message() MessageStore {
	return NewSMessage(s.db)
}
