package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
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
	var req api.FollowActionReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	if req.ActionType != 1 && req.ActionType != 2 {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: "actionType未知"})
		return
	}
	username := ctx.GetString(token.Config.IdentityKey)
	//biz
	err := c.b.Follow().Action(ctx, username, uint(req.ToUserID), uint(req.ActionType))
	if err != nil {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//操作成功
	if req.ActionType == 1 {
		ctx.JSON(200, api.FollowActionRsp{StatusCode: 0, StatusMsg: "关注成功"})
		return
	}
	ctx.JSON(200, api.FollowActionRsp{StatusCode: 0, StatusMsg: "取消关注成功"})
	return

}

func (c *CRelation) FollowList(ctx *gin.Context) {
	var req api.FollowListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FollowList(ctx, uint(req.UserID))
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
	var req api.FollowerListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.FollowerListRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FollowerList(ctx, uint(req.UserID))
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
	var req api.FriendListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.FriendListRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FriendList(ctx, uint(req.UserID))
	if err != nil {
		ctx.JSON(200, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(200, rsp)
	return
}
