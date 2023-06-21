package TikTokk

import (
	"TikTokk/internal/TikTokk/biz/video"
	"TikTokk/internal/TikTokk/store"
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

func init() {
	Config()
	Mysql()
	Jwt()
	Store()
	FeedLenInit()
	SavePath()
}

func Config() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Init_Config Success")
}

func Mysql() {
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
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	fmt.Println("Init_Mysql success")
}

func Jwt() {
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

}

func Store() {
	_ = store.NewStore(DB)
}

func FeedLenInit() {
	video.FeedLen = viper.GetInt("feed.len")

}

func SavePath() {
	video.FileSavePath = viper.GetString("fileSave.file")
	video.UploadsSavePath = viper.GetString("fileSave.uploads")
}

func InitConfigTest() {
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	viper.SetConfigName("config")
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
