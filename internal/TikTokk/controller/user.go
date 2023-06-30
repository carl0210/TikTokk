package controller

import (
	"TikTokk/api"
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(ctx.Request.URL.Query())
		fmt.Println(err)
		ctx.JSON(http.StatusOK, api.LoginUserRespond{StatusCode: 1, StatusMsg: "invalid field"})
		return
	}
	rsp, err := c.bz.Users().Login(ctx, &req)
	if err != nil {
		rsp.StatusMsg = err.Error()
		rsp.StatusCode = 1
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "登录成功!"
	rsp.Token = token.Sign(req.Username)
	ctx.JSON(http.StatusOK, rsp)
	return
}

func (c *CUser) Register(ctx *gin.Context) {
	var req api.RegisterUserRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.RegisterUserRespond{StatusCode: 1, StatusMsg: "invalid field"})
		return
	}
	rsp, err := c.bz.Users().Register(ctx, &req)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "注册成功"
	rsp.Token = token.Sign(req.Username)
	ctx.JSON(http.StatusOK, rsp)
	return
}

func (c *CUser) GetDetail(ctx *gin.Context) {
	var req api.GetDetailUserRequest
	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, api.GetDetailUserRespond{StatusCode: 1, StatusMsg: "invalid field"})
	}
	rsp, err := c.bz.Users().GetDetail(ctx, &req)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		ctx.JSON(http.StatusOK, rsp)
		return
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功!"
	ctx.JSON(http.StatusOK, rsp)
	return
}
