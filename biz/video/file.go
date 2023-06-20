package video

import (
	"TikTokk/utils"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func sentVideo(f *multipart.FileHeader, videoFileName string) error {
	//创建表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("data", videoFileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//创建文件reader
	ff, _ := f.Open()
	defer ff.Close()
	//写入表单
	if _, err := io.Copy(formFile, ff); err != nil {
		fmt.Println(err)
		return err
	}
	writer.Close()
	//发送请求
	r, err := http.Post(utils.UploadsSavePath, writer.FormDataContentType(), body)
	fmt.Println("r= ", r)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	//检查请求是否成功
	//从body中读出数据
	rsp, err := utils.FileToRsp(r.Body)
	if err != nil {
		return err
	}
	if rsp.StatusCode != 0 {
		return fmt.Errorf(rsp.StatusMsg)
	}
	return nil
}
