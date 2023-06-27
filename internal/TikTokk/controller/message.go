package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"time"
)

type IMessage interface {
	Action(ctx *gin.Context)
	Chat(ctx *gin.Context)
}

type CMessage struct {
	b biz.IBiz
}

var _ IMessage = (*CMessage)(nil)

func NewCMessage(ds store.DataStore) *CMessage {
	return &CMessage{b: biz.NewBiz(ds)}
}

func (c *CMessage) Action(ctx *gin.Context) {
	var req api.MessageActionReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.MessageActionRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//从token中获取name
	name := ctx.GetString(token.Config.IdentityKey)
	//biz
	err := c.b.Message().Action(ctx, name, req.Content, req.ToUserID)
	if err != nil {
		ctx.JSON(200, api.MessageActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(200, api.MessageActionRsp{StatusCode: 0, StatusMsg: "发送成功!"})
	return
}

func (c *CMessage) Chat(ctx *gin.Context) {
	var req api.MessageChatReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.MessageChatRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//得到name
	name := ctx.GetString(token.Config.IdentityKey)

	if req.PreMsgTime == 0 {
		req.PreMsgTime = time.Now().Add(-time.Hour).Unix()
	}
	//biz
	l, err := c.b.Message().Chat(ctx, name, req.ToUserID, req.PreMsgTime)
	if err != nil {
		ctx.JSON(200, api.MessageChatRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(200, api.MessageChatRsp{StatusCode: 0, StatusMsg: "获取成功", MessageList: l})
	return
}
