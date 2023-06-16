package controller

import (
	"TikTokk/api"
	"TikTokk/biz"
	"TikTokk/store"
	"TikTokk/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	//获取token中的用户名
	var name string
	token := ctx.Query("token")
	if len(token) != 0 {
		name, _ = utils.Parse(token, utils.Config.Key)
	} else {
		name = ""
	}

	//获取latest_time
	var latestTime time.Time
	latestTimeStr := ctx.Query("latest_time")
	//如果未传入,则为当前时间
	if latestTimeStr == "" {
		latestTime = time.Now()
	} else {
		lTime, err := strconv.Atoi(latestTimeStr)
		if err != nil {
			ctx.JSON(200, api.VideoFeedListRsp{StatusCode: 1, StatusMsg: err.Error(), NextTime: time.Now().Unix()})
			return
		}
		latestTime = time.Unix(int64(lTime), 0)
	}
	//如果latest_time比当前时间大则代表不合法
	if !latestTime.Before(time.Now()) {
		latestTime = time.Now()
	}
	rsp, err := c.b.Videos().GetVideoFeedList(ctx, name, latestTime)
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功!"
	ctx.JSON(200, rsp)
	fmt.Println("rsp=", rsp, "\tnextTime=", time.Unix(rsp.NextTime, 0).String(), "\tlatestTime=", latestTime.String())
	return
}

func (c *CVideo) PublishAction(ctx *gin.Context) {
	var rsp api.VideoPublishActionRsp
	file, err := ctx.FormFile("data")
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(200, rsp)
		return
	}
	title := ctx.PostForm("title")
	username := ctx.GetString("username")
	err = c.b.Videos().PublishAction(ctx, file, title, username)
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "上传成功"
	ctx.JSON(200, rsp)
	return

}

func (c *CVideo) PublishList(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	//如果用户ID不正确
	if err != nil || 0 > userID {
		rsp := api.VideoPublishListRsp{VideoList: nil, StatusCode: 1, StatusMsg: "用户ID不正确"}
		ctx.JSON(200, rsp)
		return
	}
	rsp, err := c.b.Videos().PublishList(ctx, userID)
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
