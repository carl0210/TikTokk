package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IRelFavorite interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type CRelFavorite struct {
	b biz.IBiz
}

var _ IRelFavorite = (*CRelFavorite)(nil)

func NewCRelFavorite(db store.DataStore) *CRelFavorite {
	return &CRelFavorite{b: biz.NewBiz(db)}
}

func (c *CRelFavorite) List(ctx *gin.Context) {
	//获取userID
	userIDStr := ctx.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusOK, api.FavoriteListRsp{
			StatusMsg:  err.Error(),
			StatusCode: 1,
		})
		return
	}
	//进入biz层
	rsp, err := c.b.FavoriteRel().List(ctx, uint(userID))
	if err != nil {
		ctx.JSON(http.StatusOK, api.FavoriteListRsp{
			StatusMsg:  err.Error(),
			StatusCode: 1,
		})
		return
	}
	rsp.StatusMsg = "获取成功"
	rsp.StatusCode = 0
	ctx.JSON(http.StatusOK, rsp)
	return

}

func (c *CRelFavorite) Action(ctx *gin.Context) {
	var req api.FavoriteActionReq
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.FavoriteActionRsp{StatusCode: 1, StatusMsg: "invalid field"})
		return
	}
	//得到token的username
	username := ctx.GetString(token.Config.IdentityKey)
	err := c.b.FavoriteRel().Action(ctx, uint(req.VideoID), uint(req.ActionType), username)
	if err != nil {
		ctx.JSON(http.StatusOK, api.FavoriteActionRsp{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, api.FavoriteActionRsp{StatusCode: 0, StatusMsg: "操作成功!"})
	return
}
