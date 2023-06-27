package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

type IComment interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type CComment struct {
	b biz.IBiz
}

var _ IComment = (*CComment)(nil)

func NewCComment(s store.DataStore) *CComment {
	return &CComment{b: biz.NewBiz(s)}
}

func (c CComment) Action(ctx *gin.Context) {
	var req api.CommentActionReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.CommentActionRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	username := ctx.GetString(token.Config.IdentityKey)
	//对不同类型进行处理
	if req.ActionType == "1" {
		//biz
		rsp, err := c.b.Comment().Create(ctx, uint(req.VideoID), username, req.CommentText)
		if err != nil {
			ctx.JSON(200, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		//创建成功
		rsp.StatusCode = 0
		rsp.StatusMsg = "创建成功"
		ctx.JSON(200, rsp)
		return

	} else if req.ActionType == "2" {
		//biz
		err := c.b.Comment().Delete(ctx, uint(req.CommentID), uint(req.VideoID), username)
		if err != nil {
			rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "comment_id不正确"}
			ctx.JSON(200, rsp)
			return
		}
		//删除成功
		ctx.JSON(200, api.CommentActionRsp{StatusCode: 0, StatusMsg: "删除成功"})
		return
	} else {
		//未知类型
		rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "未知action_type"}
		ctx.JSON(200, rsp)
		return
	}
	return
}

func (c CComment) List(ctx *gin.Context) {
	var req api.CommentListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.CommentListRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//biz
	rsp, err := c.b.Comment().List(ctx, uint(req.VideoID))
	if err != nil {
		rspE := api.CommentListRsp{StatusCode: 1, StatusMsg: err.Error()}
		ctx.JSON(200, rspE)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(200, rsp)
	return
}
