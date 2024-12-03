package minio

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"TikTokk/internal/pkg/Tlog"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var (
	endpoint        = ""
	accessKeyID     = ""
	secretAccessKey = ""
	useSSL          = false
	MinioClient     *minio.Client
)

func Init() {
	//read minio config  by viper
	endpoint = viper.GetString("minio.endpoint")
	accessKeyID = viper.GetString("minio.accessKeyID")
	secretAccessKey = viper.GetString("minio.secretAccessKey")
	useSSL = viper.GetBool("minio.useSSL")
	if endpoint == "" || accessKeyID == "" || secretAccessKey == "" {
		Tlog.Panicw("endpoint or accessKeyID or useSSL string is empty!", endpoint, accessKeyID, secretAccessKey)
		return
	}
	//create minio client
	mc, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(
			accessKeyID, secretAccessKey, "",
		),
		Secure: useSSL,
	})
	if err != nil {
		Tlog.Fatalw("create minio client failed ! error : ", err)
		return
	}
	MinioClient = mc
	Tlog.Debugw("create minio client successful")
}

func PutObject(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64) error {
	//check arg
	if objectName == "" || bucketName == "" {
		e := fmt.Errorf("objectname or bucketname string is empty ! objectname = %s, bucketName = %s", objectName, bucketName)
		return e
	}
	if reader == nil {
		e := fmt.Errorf("reader of minio.putobject is nil")
		return e
	}
	_, err := MinioClient.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func GetObject(ctx context.Context, bucketName, objectName string, expiry time.Duration) (playUrl string, err error) {
	//check arg
	if objectName == "" || bucketName == "" {
		e := fmt.Errorf("objectname or bucketname string is empty ! objectname = %s, bucketName = %s", objectName, bucketName)
		return "", e
	}
	//get object url
	u, err := MinioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, url.Values{})
	if err != nil {
		return "", err
	}
	ureq := u.RequestURI()
	Tlog.Debugw(u.String())
	playUrl = fmt.Sprintf("http://192.168.31.29:9000%s", ureq)
	return playUrl, nil
}
