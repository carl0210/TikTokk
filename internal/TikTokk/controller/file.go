package controller

import (
	"TikTokk/internal/TikTokk/biz"
	"TikTokk/internal/TikTokk/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IFile interface {
	Uploads(ctx *gin.Context)
}

type CFile struct {
	b biz.IBiz
}

var _ IFile = (*CFile)(nil)

type UploadsRsp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func NewFile(s store.DataStore) *CFile {
	return &CFile{b: biz.NewBiz(s)}
}

func (C CFile) Uploads(ctx *gin.Context) {
	//得到上传的文件
	data, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, UploadsRsp{StatusCode: 1, StatusMsg: "文件获取失败"})
		return
	}
	//保存到当地
	uploadsPath := "./asset/video/"
	ctx.SaveUploadedFile(data, uploadsPath+data.Filename)
	if err != nil {
		ctx.JSON(http.StatusOK, UploadsRsp{StatusCode: 1, StatusMsg: "文件保存失败"})
		return
	}
	ctx.JSON(http.StatusOK, UploadsRsp{StatusCode: 0, StatusMsg: "保存成功！"})
	return
}
