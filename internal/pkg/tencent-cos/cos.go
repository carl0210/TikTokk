package tencent_cos

import (
	extract_video "TikTokk/tools/t-ffmpeg"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
)

func getClient() *cos.Client {
	url, err := url.Parse("https://tiktokk-1331222828.cos.ap-guangzhou.myqcloud.com")
	if err != nil {
		panic(err)
	}
	b := &cos.BaseURL{
		BucketURL: url,
	}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "xxx",
			SecretKey: "xxx",
		},
	})
}

func UploadVideo(ctx context.Context, key string, fd *os.File) (videoPlayUrl string, err error) {
	// 上传文件
	fileStat, err := fd.Stat()
	if err != nil {
		log.Println(err)
		return "", err
	}
	opt := &cos.ObjectPutOptions{ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
		ContentLength: fileStat.Size(),
	}}
	cli := getClient()
	k := fmt.Sprintf("video/%s", key)
	_, err = cli.Object.PutFromFile(ctx, k, fd.Name(), opt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://tiktokk-1331222828.cos.ap-guangzhou.myqcloud.com/%s", k), nil
}

func UploadVideoCoverByHeader(ctx context.Context, key string, f *os.File) (photoUrl string, err error) {
	// 抽取视频的第一帧作为封面
	err = extract_video.GetVideoCover(key, f)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// 上传文件
	cli := getClient()
	k := fmt.Sprintf("photo/%s", key)
	photoFd, err := os.Open("./asset/cover/" + key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer os.Remove("./asset/cover/" + key)
	defer photoFd.Close()
	_, err = cli.Object.Put(ctx, k, photoFd, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://tiktokk-1331222828.cos.ap-guangzhou.myqcloud.com/%s", k), nil
}
