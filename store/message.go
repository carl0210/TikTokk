package store

import (
	"TikTokk/model"
	"context"
	"gorm.io/gorm"
)

type MessageStore interface {
	Create(ctx context.Context, s *model.Chat_Message) error
	List(ctx context.Context, name string, id uint, preMsgTime int64) ([]model.Chat_Message, error)
}

type SMessage struct {
	ds *gorm.DB
}

var _ MessageStore = (*SMessage)(nil)

func NewSMessage(db *gorm.DB) *SMessage {
	return &SMessage{ds: db}
}

func (s *SMessage) Create(ctx context.Context, new *model.Chat_Message) error {
	return s.ds.Create(new).Error

}
func (s *SMessage) List(ctx context.Context, name string, id uint, preMsgTime int64) ([]model.Chat_Message, error) {
	var c []model.Chat_Message
	err := s.ds.Where("create_time > ?", preMsgTime).
		Where(
			s.ds.Where(model.Chat_Message{FromUserID: id, ToUserName: name}).
				Or(model.Chat_Message{FromUserName: name, ToUserID: id}),
		).
		Order("create_time").
		Find(&c).Error
	return c, err
}
