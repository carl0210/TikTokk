package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.VideoFeedListRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	//获取token中的用户名
	var name string
	t := ctx.Query("token")
	if len(t) != 0 {
		name, _ = token.Parse(t, token.Config.IdentityKey)
	} else {
		name = ""
	}

	//如果未传入,则为当前时间
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}
	//如果latest_time比当前时间大则代表不合法
	if req.LatestTime > time.Now().Unix() {
		req.LatestTime = time.Now().Unix()
	}
	rsp, err := c.b.Videos().GetVideoFeedList(ctx, name, req.LatestTime)
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功!"
	ctx.JSON(200, rsp)
	return
}

func (c *CVideo) PublishAction(ctx *gin.Context) {
	var req api.VideoPublishActionReq
	if err := ctx.BindWith(&req, binding.Form); err != nil {
		ctx.JSON(200, api.VideoPublishActionRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	username := ctx.GetString(token.Config.IdentityKey)
	err := c.b.Videos().PublishAction(ctx, req.File, req.Title, username)
	if err != nil {
		ctx.JSON(200, api.VideoPublishActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	ctx.JSON(200, api.VideoPublishActionRsp{StatusCode: 0, StatusMsg: "上传成功"})
	return

}

func (c *CVideo) PublishList(ctx *gin.Context) {
	var req api.VideoPublishListReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(200, api.VideoPublishListRsp{StatusCode: 1, StatusMsg: "invalid filed"})
		return
	}
	rsp, err := c.b.Videos().PublishList(ctx, int(req.UserID))
	if err != nil {
		rsp := api.VideoPublishListRsp{VideoList: nil, StatusCode: 1, StatusMsg: err.Error()}
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	ctx.JSON(200, rsp)
	return
}
