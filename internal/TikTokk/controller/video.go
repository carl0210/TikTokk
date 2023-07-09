package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

type IVideo interface {
	Feed(ctx *gin.Context)
	PublishAction(ctx *gin.Context)
	PublishList(ctx *gin.Context)
}

type CVideo struct {
	b biz.IBiz
}

var _ IVideo = (*CVideo)(nil)

func NewCVideo(s store.DataStore) *CVideo {
	return &CVideo{b: biz.NewBiz(s)}
}

func (c *CVideo) Feed(ctx *gin.Context) {
	var req api.VideoFeedListReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, api.VideoFeedListRsp{StatusCode: 1, StatusMsg: "Feed invalid field"})
		return
	}
	//获取token中的用户ID
	var userID int
	userIDStr, err := token.Parse(req.Token, token.Config.Key)
	if err != nil {
		userID = 0
	} else {

		userID, err = strconv.Atoi(userIDStr)
		if err != nil {
			userID = 0
		}

	}

	//如果未传入或如果latest_time比当前时间大则代表不合法,则为当前时间
	if req.LatestTime == 0 || req.LatestTime > time.Now().Unix() {
		req.LatestTime = time.Now().Unix()
	}

	rsp, err := c.b.Videos().GetVideoFeedList(ctx, uint(userID), req.LatestTime)
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(http.StatusOK, rsp)
		return
	}

	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功!"
	ctx.JSON(http.StatusOK, rsp)
	return
}

func (c *CVideo) PublishAction(ctx *gin.Context) {
	var req api.VideoPublishActionReq
	if err := ctx.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusOK, api.VideoPublishActionRsp{StatusCode: 1, StatusMsg: "VideoPublishAction invalid field"})
		return
	}
	userID, err := GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, api.CommentActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	err = c.b.Videos().PublishAction(ctx, req.Data, req.Title, userID)
	if err != nil {
		ctx.JSON(http.StatusOK, api.VideoPublishActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, api.VideoPublishActionRsp{StatusCode: 0, StatusMsg: "上传成功"})
	return

}

func (c *CVideo) PublishList(ctx *gin.Context) {
	var req api.VideoPublishListReq
	if err := ctx.ShouldBindQuery(&req); err != nil || req.UserID < 0 {
		ctx.JSON(http.StatusOK, api.VideoPublishListRsp{StatusCode: 1, StatusMsg: "VideoPublishList invalid field"})
		return
	}
	rsp, err := c.b.Videos().PublishList(ctx, uint(req.UserID))
	if err != nil {
		ctx.JSON(http.StatusOK, api.VideoPublishListRsp{VideoList: nil, StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, rsp)
	return
}
