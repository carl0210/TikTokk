package t_ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"io"
	"mime/multipart"
	"os"
	// ffmpeg "github.com/u2takey/ffmpeg-go"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
)

// 从
func GetVideoCover(photoName string, f *os.File) (err error) {
	// 使用ffmepg抽取第一帧
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(f.Name()).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg", "b:v": "128k"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Println("生成缩略图失败 - ffmpeg，err = ", err)
		return err
	}
	// 使用imaging 将第一帧转化为图片
	image, err := imaging.Decode(buf)
	if err != nil {
		log.Println("生成缩略图失败 - 转化图片，err = ", err)
		return err
	}
	image = imaging.Fit(image, 720, 1440, imaging.NearestNeighbor)
	err = imaging.Save(image, "./asset/cover/"+photoName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
	// ffmpeg.Input()
}

// 向临时文件写入压缩后的视频
func CompressVideo(src multipart.File, dstFileName string) error {
	// 将src中的视频写入另一临时文件a
	tempFile, err := os.CreateTemp("./tmp", "tmp.pre.*.mp4")
	if err != nil {
		log.Println(err)
		return err
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	_, err = io.Copy(tempFile, src)
	if err != nil {
		log.Println(err)
		return err
	}
	// 对使用ffmpeg对a进行压缩，并将压缩后的视频写入dstFileName
	err = ffmpeg.Input(tempFile.Name()).Output(dstFileName, ffmpeg.KwArgs{"b:v": "8M"}).OverWriteOutput().
		Run()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
