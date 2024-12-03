package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/Tlog"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := ctx.ShouldBindQuery(&req); err != nil || req.ToUserID < 0 || req.ActionType != 1 && req.ActionType != 2 {
		ctx.JSON(http.StatusOK, api.FollowActionRsp{StatusCode: 1, StatusMsg: "FollowAction invalid field"})
		return
	}

	userID, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//biz
	err = c.b.Follow().Action(ctx, userID, uint(req.ToUserID), uint(req.ActionType))
	if err != nil {
		ctx.JSON(http.StatusOK, api.FollowActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	//操作成功
	if req.ActionType == 1 {
		ctx.JSON(http.StatusOK, api.FollowActionRsp{StatusCode: 0, StatusMsg: "关注成功"})
		return
	}
	ctx.JSON(http.StatusOK, api.FollowActionRsp{StatusCode: 0, StatusMsg: "取消关注成功"})
	return

}

func (c *CRelation) FollowList(ctx *gin.Context) {
	var req api.FollowListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.FollowListRsp{StatusCode: 1, StatusMsg: "FollowList invalid field"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FollowList(ctx, uint(req.UserID))
	if err != nil {
		ctx.JSON(http.StatusOK, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, rsp)
	return
}

func (c *CRelation) FollowerList(ctx *gin.Context) {
	var req api.FollowerListReq
	if err := ctx.ShouldBindQuery(&req); err != nil || req.UserID < 0 {
		ctx.JSON(http.StatusOK, api.FollowerListRsp{StatusCode: 1, StatusMsg: "FollowerList invalid field"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FollowerList(ctx, uint(req.UserID))
	if err != nil {
		Tlog.Infow(err.Error())
		ctx.JSON(http.StatusOK, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, rsp)
	return
}

func (c *CRelation) FriendListList(ctx *gin.Context) {
	var req api.FriendListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.FriendListRsp{StatusCode: 1, StatusMsg: "FriendList invalid field"})
		return
	}
	//biz
	rsp, err := c.b.Follow().FriendList(ctx, uint(req.UserID))
	if err != nil {
		ctx.JSON(http.StatusOK, api.FollowListRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, rsp)
	return
}
