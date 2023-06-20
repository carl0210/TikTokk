package controller

import (
	"TikTokk/api"
	"TikTokk/biz"
	"TikTokk/store"
	"TikTokk/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IRelFavorite interface {
	FollowAction(ctx *gin.Context)
	FollowList(ctx *gin.Context)
}

type CRelFavorite struct {
	b biz.IBiz
}

var _ IRelFavorite = (*CRelFavorite)(nil)

func (c *CRelFavorite) FollowList(ctx *gin.Context) {
	//获取userID
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(200, api.FavoriteListRsp{
			StatusMsg:  err.Error(),
			StatusCode: 1,
		})
		return
	}
	//进入biz层
	rsp, err := c.b.FavoriteRel().List(ctx, uint(userID))
	if err != nil {
		ctx.JSON(200, api.FavoriteListRsp{
			StatusMsg:  err.Error(),
			StatusCode: 1,
		})
		return
	}
	rsp.StatusMsg = "获取成功"
	rsp.StatusCode = 0
	ctx.JSON(200, rsp)
	return

}

func NewCRelFavorite(db store.DataStore) *CRelFavorite {
	return &CRelFavorite{b: biz.NewBiz(db)}
}

func (c *CRelFavorite) FollowAction(ctx *gin.Context) {
	var rsp api.FavoriteActionRsp
	//将参数转为整型
	//videoID
	videoIDStr := ctx.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(200, rsp)
		return
	}
	//actionType
	actionTypeStr := ctx.Query("action_type")
	actionType, err := strconv.Atoi(actionTypeStr)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(200, rsp)
		return
	}
	//得到token的username
	username := ctx.GetString(utils.Config.IdentityKey)
	if len(username) == 0 {
		rsp.StatusCode = 1
		rsp.StatusMsg = "用户名为空"
		ctx.JSON(200, rsp)
		return
	}
	err = c.b.FavoriteRel().Action(ctx, uint(videoID), uint(actionType), username)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(200, rsp)
		return
	}

	rsp.StatusCode = 0
	rsp.StatusMsg = "操作成功!"
	ctx.JSON(200, rsp)
	return
}
