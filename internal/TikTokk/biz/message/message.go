package message

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/tools"
	"context"
	"time"
)

type MessageBiz interface {
	Action(ctx context.Context, name, content string, toUserID int64) error
	Chat(ctx context.Context, name string, toUserID int64, preMsgTime int64) ([]api.MessageDetailRsp, error)
}

type BMessage struct {
	ds store.DataStore
}

var _ MessageBiz = (*BMessage)(nil)

func New(s store.DataStore) *BMessage {
	return &BMessage{ds: s}
}

func (b BMessage) Action(ctx context.Context, name, content string, toUserID int64) error {
	//获取发送人信息,并检验是否存在
	u, err := b.ds.Users().Get(ctx, &model.User{Name: name})
	if err != nil {
		return err
	}
	//获取接受人信息,并检验是否存在
	t, err := b.ds.Users().Get(ctx, &model.User{UserId: uint(toUserID)})
	if err != nil {
		return err
	}
	//创建记录
	err = b.ds.Message().Create(ctx, &model.Chat_Message{
		FromUserID:   u.UserId,
		FromUserName: u.Name,
		ToUserID:     t.UserId,
		ToUserName:   t.Name,
		Content:      content,
		CreateTime:   time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	//创建成功
	return nil
}

func (b BMessage) Chat(ctx context.Context, name string, toUserID int64, preMsgTime int64) ([]api.MessageDetailRsp, error) {
	//获取列表
	list, err := b.ds.Message().List(ctx, name, uint(toUserID), preMsgTime)
	if err != nil {
		return nil, err
	}
	//转化
	rsp := tools.MessagestoRsp(list)
	return rsp, nil
}
