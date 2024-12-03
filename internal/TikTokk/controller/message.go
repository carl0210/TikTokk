package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := ctx.ShouldBindQuery(&req); err != nil || req.ToUserID < 0 || req.ActionType != 1 {
		ctx.JSON(http.StatusOK, api.MessageActionRsp{StatusCode: 1, StatusMsg: "MessageAction invalid field"})
		return
	}
	//从token中获取id
	userID, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//biz
	err = c.b.Message().Action(ctx, req.Content, userID, uint(req.ToUserID))
	if err != nil {
		ctx.JSON(http.StatusOK, api.MessageActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.MessageActionRsp{StatusCode: 0, StatusMsg: "发送成功!"})
	return
}

func (c *CMessage) Chat(ctx *gin.Context) {
	var req api.MessageChatReq
	if err := ctx.ShouldBindQuery(&req); err != nil || req.ToUserID < 0 {
		ctx.JSON(http.StatusOK, api.MessageChatRsp{StatusCode: 1, StatusMsg: "MessageChat invalid field"})
		return
	}
	//得到name
	userID, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	if req.PreMsgTime == 0 {
		req.PreMsgTime = time.Now().Add(-time.Hour).Unix()
	}
	//biz
	l, err := c.b.Message().Chat(ctx, userID, uint(req.ToUserID), req.PreMsgTime)
	if err != nil {
		ctx.JSON(http.StatusOK, api.MessageChatRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.MessageChatRsp{StatusCode: 0, StatusMsg: "获取成功", MessageList: l})
	return
}
