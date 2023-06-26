package TikTokk

import (
	"TikTokk/internal/TikTokk/biz/video"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/Tlog"
	"TikTokk/internal/pkg/token"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var once sync.Once

func TikTokInit() {
	Config()
	Logg()
	Mysql()
	TikTokk()
	Tlog.Infow("TikTokInit Successful")
}

func Logg() {
	Tlog.Init(Tlog.LogOption())
	Tlog.Infow("init_logg successful")
	defer Tlog.Sync()
}

func Config() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Config Init error =", err.Error())
		return
	}
	log.Println("Config Init successful")
}

func Mysql() {
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:      logger.LogLevel(viper.GetInt("mysql.Tlog-level")),
			SlowThreshold: time.Second,
			Colorful:      true,
		})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.ip"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.config"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		Tlog.Panicw("Init_mysql error =", err.Error())
		return
	}
	DB = db

	Tlog.Infow("Init_Mysql successful")
}

// TikTokk 初始化业务代码
func TikTokk() {
	//初始化token
	once.Do(func() {
		key := viper.GetString("jwt.key")
		identity := viper.GetString("jwt.identityKey")

		if key != "" {
			token.Config.Key = key
		}
		if identity != "" {
			token.Config.IdentityKey = identity
		}

	})
	//初始化store层
	_ = store.NewStore(DB)
	//初始化视频biz层
	video.FeedLen = viper.GetInt("feed.len")
	video.FileSavePath = viper.GetString("fileSave.file")
	video.UploadsSavePath = viper.GetString("fileSave.uploads")

	Tlog.Infow("Init_TikTokk successful")
}
