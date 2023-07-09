package store

import (
	"TikTokk/internal/TikTokk/model"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type MessageStore interface {
	Create(ctx context.Context, s *model.ChatMessage) error
	List(ctx context.Context, userID, toUserID uint, preMsgTime int64) ([]model.ChatMessage, error)
}

type SMessage struct {
	ds *gorm.DB
	rc *redis.Client
}

var _ MessageStore = (*SMessage)(nil)

func NewSMessage(db *gorm.DB, rc *redis.Client) *SMessage {
	return &SMessage{ds: db, rc: rc}
}

func (s *SMessage) Create(ctx context.Context, new *model.ChatMessage) error {
	return s.ds.Create(new).Error
}
func (s *SMessage) List(ctx context.Context, userID, toUserID uint, preMsgTime int64) ([]model.ChatMessage, error) {
	var c []model.ChatMessage
	err := s.ds.Where("create_time > ?", preMsgTime).
		Where(
			s.ds.Where(model.ChatMessage{FromUserID: userID, ToUserID: toUserID}).
				Or(model.ChatMessage{FromUserID: toUserID, ToUserID: userID}),
		).
		Order("create_time").
		Find(&c).Error
	return c, err
}
