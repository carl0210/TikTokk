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
	FileSavePath string
	DB           *gorm.DB
	once         sync.Once
	FeedLen      int
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
			LogLevel:      logger.Info,
			SlowThreshold: time.Second,
			Colorful:      true,
		})
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: log})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	fmt.Println("Init_Mysql success")
}

func InitJwt(key, identity string) {
	once.Do(func() {
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

func InitFileSavePath() {
	FileSavePath = viper.GetString("fileSave.path")
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
