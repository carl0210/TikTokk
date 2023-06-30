package TikTokk

import (
	"TikTokk/internal/TikTokk/biz/video"
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/TikTokk/store"
	"TikTokk/internal/pkg/Tlog"
	"TikTokk/internal/pkg/token"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
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
	master := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		viper.GetString("mysql-master.user"),
		viper.GetString("mysql-master.password"),
		viper.GetString("mysql-master.ip"),
		viper.GetInt("mysql-master.port"),
		viper.GetString("mysql-master.database"),
		viper.GetString("mysql-master.config"),
	)
	slave1 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		viper.GetString("mysql-slave-1.user"),
		viper.GetString("mysql-slave-1.password"),
		viper.GetString("mysql-slave-1.ip"),
		viper.GetInt("mysql-slave-1.port"),
		viper.GetString("mysql-slave-1.database"),
		viper.GetString("mysql-slave-1.config"),
	)
	slave2 := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		viper.GetString("mysql-slave-2.user"),
		viper.GetString("mysql-slave-2.password"),
		viper.GetString("mysql-slave-2.ip"),
		viper.GetInt("mysql-slave-2.port"),
		viper.GetString("mysql-slave-2.database"),
		viper.GetString("mysql-slave-2.config"),
	)
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true, // Don't include params in the SQL log
			LogLevel:                  logger.LogLevel(viper.GetInt("mysql.log-level")),
			SlowThreshold:             time.Second,
			Colorful:                  true,
		})

	db, err := gorm.Open(mysql.Open(master), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		Tlog.Panicw("Init_mysql error =", err.Error())
		return
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(slave1), mysql.Open(slave2)},
		Policy:   dbresolver.RandomPolicy{},
	}, &model.User{}, &model.UserFollowed{}, &model.Chat_Message{}, &model.Comment{}, &model.UserFavorite{}, &model.Video{}))

	DB = db
	log.Info(context.Background(), "Init_Mysql successful")
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
