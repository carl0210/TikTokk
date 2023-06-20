package utils

import (
	"TikTokk/store"
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

var (
	FileSavePath    string
	DB              *gorm.DB
	once            sync.Once
	FeedLen         int
	UploadsSavePath string
)

func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Init_Config Success")
}

func InitMysql() {
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:      logger.LogLevel(viper.GetInt("mysql.log-level")),
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
	fmt.Println(dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	fmt.Println("Init_Mysql success")
}

func InitJwt() {
	once.Do(func() {
		key := viper.GetString("jwt.key")
		identity := viper.GetString("jwt.identityKey")
		if key != "" {
			Config.Key = key
		}
		if identity != "" {
			Config.IdentityKey = identity
		}
	})

}

func InitStore() {
	_ = store.NewStore(DB)
}

func InitFeedlen() {
	FeedLen = viper.GetInt("feed.len")

}

func InitSavePath() {
	FileSavePath = viper.GetString("fileSave.file")
	UploadsSavePath = viper.GetString("fileSave.uploads")
}

func InitConfigTest() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitMysqlTest() {
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
