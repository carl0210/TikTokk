package controller

import (
	"TikTokk/api"
	"TikTokk/biz"
	"TikTokk/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	//验证action_type
	actionType := ctx.Query("action_type")
	if actionType != "1" {
		ctx.JSON(200, api.MessageActionRsp{StatusCode: 1, StatusMsg: "action_type不正确"})
		return
	}
	//从token中获取name,query参数中得到toUserID、content
	name := ctx.GetString("username")
	toUserIDStr := ctx.Query("to_user_id")
	content := ctx.Query("content")
	//转化toUserIDStr为整型
	toUserID, err := strconv.Atoi(toUserIDStr)
	if err != nil {
		ctx.JSON(200, api.MessageActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//biz
	err = c.b.Message().Action(ctx, name, content, toUserID)
	if err != nil {
		ctx.JSON(200, api.MessageActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(200, api.MessageActionRsp{StatusCode: 0, StatusMsg: "发送成功!"})
	return
}

func (c *CMessage) Chat(ctx *gin.Context) {
	//得到name、to_user_id、pre_msg_time
	name := ctx.GetString("username")
	toUserIDStr := ctx.Query("to_user_id")
	preMsgTimeStr := ctx.Query("pre_msg_time")
	//转化toUserID、preMsgTime
	toUserID, err := strconv.Atoi(toUserIDStr)
	if err != nil {
		ctx.JSON(200, api.MessageChatRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	preMsgTime, err := strconv.ParseInt(preMsgTimeStr, 10, 64)
	if err != nil {
		ctx.JSON(200, api.MessageChatRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if preMsgTime == 0 {
		preMsgTime = time.Now().Add(-time.Hour).Unix()
	}
	//biz
	l, err := c.b.Message().Chat(ctx, name, toUserID, preMsgTime)
	if err != nil {
		ctx.JSON(200, api.MessageChatRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	fmt.Println(l)
	ctx.JSON(200, api.MessageChatRsp{StatusCode: 0, StatusMsg: "获取成功", MessageList: l})
	return
}
