package store

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

type DataStore interface {
	Users() UserStore
	Videos() VideoStore
	UserFollowRelation() IUserFollowRelation
	VideoFavoriteRelation() IVideoFavoriteRelation
	Comment() CommentStore
	Message() MessageStore
}

var (
	once sync.Once
	S    *SData
)

type SData struct {
	db *gorm.DB
	rc *redis.Client
}

var _ DataStore = (*SData)(nil)

func NewStore(db *gorm.DB, rc *redis.Client) *SData {
	once.Do(func() {
		S = &SData{db: db, rc: rc}
	})
	return S
}

func (s *SData) Videos() VideoStore {
	return NewVideos(s.db, s.rc)
}

func (s *SData) Users() UserStore {
	return NewUsers(s.db, s.rc)
}

func (s *SData) UserFollowRelation() IUserFollowRelation {
	return NewSUserFollowRelation(s.db, s.rc)
}

func (s *SData) VideoFavoriteRelation() IVideoFavoriteRelation {
	return NewSVideoFavoriteRelation(s.db, s.rc)
}

func (s *SData) Comment() CommentStore {
	return NewSComment(s.db, s.rc)
}

func (s *SData) Message() MessageStore {
	return NewSMessage(s.db, s.rc)
}
