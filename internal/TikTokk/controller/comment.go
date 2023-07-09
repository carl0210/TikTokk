package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := ctx.ShouldBindQuery(&req); err != nil || req.CommentID < 0 || req.VideoID < 0 {
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: "CommentAction invalid field"})
		return
	}
	userID, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	//对不同类型进行处理
	if req.ActionType == "1" {
		//biz
		rsp, err := c.b.Comment().Create(ctx, uint(req.VideoID), userID, req.CommentText)
		if err != nil {
			ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		//创建成功
		rsp.StatusCode = 0
		rsp.StatusMsg = "创建成功"
		ctx.JSON(http.StatusOK, rsp)
		return

	} else if req.ActionType == "2" {
		//biz
		err := c.b.Comment().Delete(ctx, uint(req.CommentID), uint(req.VideoID), userID)
		if err != nil {
			rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "comment_id不正确"}
			ctx.JSON(http.StatusOK, rsp)
			return
		}
		//删除成功
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 0, StatusMsg: "删除成功"})
		return
	} else {
		//未知类型
		rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "未知action_type"}
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	return
}

func (c CComment) List(ctx *gin.Context) {
	var req api.CommentListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.CommentListRsp{StatusCode: 1, StatusMsg: "CommentList invalid field"})
		return
	}
	//biz
	rsp, err := c.b.Comment().List(ctx, uint(req.VideoID))
	if err != nil {
		rspE := api.CommentListRsp{StatusCode: 1, StatusMsg: err.Error()}
		ctx.JSON(http.StatusOK, rspE)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, rsp)
	return
}
