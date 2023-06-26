package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/Tlog"
	"TikTokk/internal/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IRelation interface {
	FollowAction(ctx *gin.Context)
	FollowList(ctx *gin.Context)
	FollowerList(ctx *gin.Context)
	FriendListList(ctx *gin.Context)
}

type CRelation struct {
	b biz.IBiz
}

var _ IRelation = (*CRelation)(nil)

func NewCRelFollow(ds store.DataStore) *CRelation {
	return &CRelation{b: biz.NewBiz(ds)}
}

func (c *CRelation) FollowAction(ctx *gin.Context) {
	//得到username、action_type、to_user_id
	//得到操作类型
	Tlog.Infow("FollowAction callers", "request header=", ctx.Request.Header)
	fmt.Println("ctx.request.header=", ctx.Request.Header)
	actionTypeStr := ctx.Query("action_type")
	actionType, err := strconv.Atoi(actionTypeStr)
	if err != nil {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: "actionType不合法"})
		return
	}
	if actionType != 1 && actionType != 2 {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: "actionType未知"})
		return
	}
	toUserIDStr := ctx.Query("to_user_id")
	toUserID, err := strconv.Atoi(toUserIDStr)
	if err != nil {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: "to_user_id不合法"})
		return
	}
	username := ctx.GetString(token.Config.IdentityKey)
	//biz
	err = c.b.Follow().Action(ctx, username, uint(toUserID), uint(actionType))
	if err != nil {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//操作成功
	if actionType == 1 {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 0, StatusMsg: "关注成功"})
		return
	}
	ctx.JSON(200, api.FollowActionRsp{StatusCode: 0, StatusMsg: "取消关注成功"})
	return

}

func (c *CRelation) FollowList(ctx *gin.Context) {
	//得到user_id
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: "userID不合法"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FollowList(ctx, uint(userID))
	if err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(200, rsp)
	return
}

func (c *CRelation) FollowerList(ctx *gin.Context) {
	//得到user_id
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: "userID不合法"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FollowerList(ctx, uint(userID))
	if err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(200, rsp)
	return
}

func (c *CRelation) FriendListList(ctx *gin.Context) {
	//得到user_id
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(200, api.FriendListRsp{StatusCode: 0, StatusMsg: "userID不合法"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FriendList(ctx, uint(userID))
	if err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(200, rsp)
	return
}
