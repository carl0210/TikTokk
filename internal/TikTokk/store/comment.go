package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type CommentStore interface {
	Create(ctx context.Context, c *model.Comment) error
	Delete(ctx context.Context, cID uint) error
	Get(ctx context.Context, c *model.Comment) (*model.Comment, error)
	Transaction(ctx context.Context, f func(tx *gorm.DB) error) error
	ListLen(ctx context.Context, l, videoID uint) ([]model.Comment, error)
	List(ctx context.Context, videoID uint) ([]model.Comment, error)
}

type SComment struct {
	db *gorm.DB
	rc *redis.Client
}

var _ CommentStore = (*SComment)(nil)

func NewSComment(db *gorm.DB, rc *redis.Client) *SComment {
	return &SComment{db: db, rc: rc}
}

func (s SComment) Get(ctx context.Context, c *model.Comment) (*model.Comment, error) {
	var r model.Comment
	err := s.db.Where(c).First(&r).Error
	return &r, err
}

func (s SComment) Create(ctx context.Context, c *model.Comment) error {
	return s.db.Create(c).Error
}

func (s SComment) Delete(ctx context.Context, cID uint) error {
	return s.db.Delete(&model.Comment{CommentID: cID}).Error
}

func (s SComment) Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	return s.db.Transaction(f)
}

// List 倒序
func (s SComment) ListLen(ctx context.Context, l, videoID uint) ([]model.Comment, error) {
	rsp := make([]model.Comment, l)
	err := s.db.Table("comments").Order("create_date desc").Where("video_id", videoID).Find(&rsp).Error
	return rsp, err
}

func (s SComment) List(ctx context.Context, videoID uint) ([]model.Comment, error) {
	var rsp []model.Comment
	err := s.db.Table("comments").Order("create_date desc").Where("video_id", videoID).Find(&rsp).Error
	return rsp, err
}
