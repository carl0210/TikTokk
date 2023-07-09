package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type IVideoFavoriteRelation interface {
	ListLen(ctx context.Context, userID uint, len int64) ([]model.Video, error)
	Transaction(ctx context.Context, f func(db *gorm.DB) error) error
	Get(ctx context.Context, VideoId, userId uint) (new *model.UserFavorite, err error)
	Create(ctx context.Context, favorite *model.UserFavorite) error
	Update(ctx context.Context, VideoId uint, username string, isFavorite bool) error
	FirstOrCreate(ctx context.Context, VideoId, userId uint, userName string) (new *model.UserFavorite, err error)
}

type SVideoFavoriteRelation struct {
	db *gorm.DB
	rc *redis.Client
}

var _ IVideoFavoriteRelation = (*SVideoFavoriteRelation)(nil)

func (s *SVideoFavoriteRelation) ListLen(ctx context.Context, userID uint, len int64) ([]model.Video, error) {
	l := make([]model.Video, len)
	err := s.db.Where("video_id IN (?)",
		s.db.Table("user_favorites").Select("video_id").Where("user_id= ?  AND is_favorite = ? ", userID, 1)).Find(&l).Error
	return l, err
}

func NewSVideoFavoriteRelation(db *gorm.DB, rc *redis.Client) *SVideoFavoriteRelation {
	return &SVideoFavoriteRelation{db: db, rc: rc}
}

func (s *SVideoFavoriteRelation) Get(ctx context.Context, VideoId, userId uint) (*model.UserFavorite, error) {
	var r model.UserFavorite
	err := s.db.Where("user_id = ? AND video_id = ?", userId, VideoId).First(&r).Error
	return &r, err
}

func (s *SVideoFavoriteRelation) Create(ctx context.Context, favorite *model.UserFavorite) error {
	return s.db.Create(favorite).Error
}

func (s *SVideoFavoriteRelation) FirstOrCreate(ctx context.Context, VideoId, userId uint, userName string) (*model.UserFavorite, error) {
	var r model.UserFavorite
	err := s.db.FirstOrCreate(&r, model.UserFavorite{UserId: userId, VideoId: VideoId, UserName: userName}).Error
	return &r, err
}

func (s *SVideoFavoriteRelation) Update(ctx context.Context, VideoId uint, username string, isFavorite bool) error {
	err := s.db.Model(&model.UserFavorite{VideoId: VideoId, UserName: username}).Update("is_favorite", isFavorite).Error
	return err

}

func (s *SVideoFavoriteRelation) Transaction(ctx context.Context, f func(db *gorm.DB) error) error {

	return s.db.Transaction(f)
}
