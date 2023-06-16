package controller

import (
	"TikTokk/api"
	"TikTokk/biz"
	"TikTokk/store"
	"TikTokk/utils"
	"github.com/gin-gonic/gin"
)

type IUser interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetDetail(c *gin.Context)
}

type CUser struct {
	bz biz.IBiz
}

var _ IUser = (*CUser)(nil)

func NewCUser(db store.DataStore) *CUser {
	return &CUser{bz: biz.NewBiz(db)}
}

func (c *CUser) Login(ctx *gin.Context) {
	var req api.LoginUserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	rsp, err := c.bz.Users().Login(ctx, &req)
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "登录成功!"
	rsp.Token = utils.Sign(req.Username)
	ctx.JSON(200, rsp)
	return
}

func (c *CUser) Register(ctx *gin.Context) {
	var req api.RegisterUserRequest
	req.Username = ctx.Query("username")
	req.Password = ctx.Query("password")
	rsp, err := c.bz.Users().Register(ctx, &req)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "注册成功"
	rsp.Token = utils.Sign(req.Username)
	ctx.JSON(200, rsp)
	return
}

func (c *CUser) GetDetail(ctx *gin.Context) {
	var req api.GetDetailUserRequest
	req.UserID = ctx.Query("user_id")
	rsp, err := c.bz.Users().GetDetail(ctx, &req)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(200, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功!"
	ctx.JSON(200, rsp)
	return
}
