package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IComment interface {
	FollowAction(ctx *gin.Context)
	FollowList(ctx *gin.Context)
}

type CComment struct {
	b biz.IBiz
}

var _ IComment = (*CComment)(nil)

func NewCComment(s store.DataStore) *CComment {
	return &CComment{b: biz.NewBiz(s)}
}

func (c CComment) FollowAction(ctx *gin.Context) {
	//得到操作类型
	opType := ctx.Query("action_type")
	//得到必带的参数
	videoIDStr := ctx.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "video_id不正确", Comment: api.CommentDetailRsp{}}
		ctx.JSON(200, rsp)
		return
	}
	username := ctx.GetString(token.Config.IdentityKey)
	//对不同类型进行处理
	if opType == "1" {
		//获取参数comment_id
		text := ctx.Query("comment_text")
		//biz
		rsp, err := c.b.Comment().Create(ctx, uint(videoID), username, text)
		if err != nil {
			rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error(), Comment: api.CommentDetailRsp{}}
			ctx.JSON(200, rsp)
			return
		}
		//创建成功
		rsp.StatusCode = 0
		rsp.StatusMsg = "创建成功"
		ctx.JSON(200, rsp)
		return

	} else if opType == "2" {
		//获取参数comment_id
		commentIDStr := ctx.Query("comment_id")
		commentID, err := strconv.Atoi(commentIDStr)
		if err != nil {
			rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "comment_id不正确", Comment: api.CommentDetailRsp{}}
			ctx.JSON(200, rsp)
			return
		}
		//biz
		err = c.b.Comment().Delete(ctx, uint(commentID), uint(videoID), username)
		if err != nil {
			rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "comment_id不正确", Comment: api.CommentDetailRsp{}}
			ctx.JSON(200, rsp)
			return
		}
		//删除成功
		rsp := api.CommentActionRsp{StatusCode: 0, StatusMsg: "删除成功", Comment: api.CommentDetailRsp{}}
		ctx.JSON(200, rsp)
		return
	} else {
		//未知类型
		rsp := api.CommentActionRsp{StatusCode: 1, StatusMsg: "未知action_type", Comment: api.CommentDetailRsp{}}
		ctx.JSON(200, rsp)
		return
	}
	return
}

func (c CComment) FollowList(ctx *gin.Context) {
	//得到video_id
	videoIDStr := ctx.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		rsp := api.CommentListRsp{StatusCode: 1, StatusMsg: "video_id不正确"}
		ctx.JSON(200, rsp)
		return
	}
	//biz
	rsp, err := c.b.Comment().List(ctx, uint(videoID))
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
