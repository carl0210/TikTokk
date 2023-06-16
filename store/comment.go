package store

import (
	"TikTokk/model"
	"context"
	"gorm.io/gorm"
)

type CommentStore interface {
	Create(ctx context.Context, c *model.Comment) error
	Delete(ctx context.Context, cID uint) error
	GetByCommentID(ctx context.Context, cID uint) (*model.Comment, error)
	GetByUName(ctx context.Context, vID uint, uName, createDate string) (*model.Comment, error)
	GetByUID(ctx context.Context, uID, vID uint, createDate string) (*model.Comment, error)
	Transaction(ctx context.Context, f func(tx *gorm.DB) error) error
	ListLen(ctx context.Context, l, videoID uint) ([]model.Comment, error)
	List(ctx context.Context, videoID uint) ([]model.Comment, error)
}

type SComment struct {
	db *gorm.DB
}

var _ CommentStore = (*SComment)(nil)

func NewSComment(db *gorm.DB) *SComment {
	return &SComment{db: db}
}

func (s SComment) GetByCommentID(ctx context.Context, cID uint) (*model.Comment, error) {
	var r model.Comment
	err := s.db.Table("comments").Where("comment_id=?", cID).First(&r).Error
	return &r, err
}

func (s SComment) GetByUID(ctx context.Context, uID, vID uint, createDate string) (*model.Comment, error) {
	var r model.Comment
	err := s.db.Table("comments").Where(
		"user_id=? AND create_date=? AND video_id=?", uID, createDate, vID,
	).First(&r).Error
	return &r, err
}

func (s SComment) GetByUName(ctx context.Context, vID uint, uName, createDate string) (*model.Comment, error) {
	var r model.Comment
	err := s.db.Table("comments").Where(
		"user_name=? AND create_date=? AND video_id=?", uName, createDate, vID,
	).First(&r).Error
	return &r, err
}
func (s SComment) Create(ctx context.Context, c *model.Comment) error {
	return s.db.Create(c).Error
}

func (s SComment) Delete(ctx context.Context, cID uint) error {
	return s.db.Delete(&model.Comment{Id: cID}).Error
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
