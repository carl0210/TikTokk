package main

import (
	"TikTokk/internal/TikTokk/model"
	"TikTokk/internal/pkg/Tlog"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"testing"
	"time"
)

var DB *gorm.DB

func TestMain(m *testing.M) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Config Init error =", err.Error())
		return
	}

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
	}, &model.User{}, &model.UserFollowed{}, &model.ChatMessage{}, &model.Comment{}, &model.UserFavorite{}, &model.Video{}))

	DB = db
}

func BenchmarkSingleDB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var u model.User
		DB.Where(model.User{UserID: 1}).First(&u)
	}

}
