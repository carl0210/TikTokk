package main

import (
	"TikTokk/internal/TikTokk"
	"TikTokk/internal/TikTokk/model"
)

func main() {
	TikTokk.Config()
	TikTokk.Mysql()
	TikTokk.DB.AutoMigrate(&model.UserFollowed{})
	TikTokk.DB.AutoMigrate(&model.User{})
	TikTokk.DB.AutoMigrate(&model.UserFavorite{})
	TikTokk.DB.AutoMigrate(&model.Comment{})
	TikTokk.DB.AutoMigrate(&model.Chat_Message{})
	TikTokk.DB.AutoMigrate(&model.Video{})

	// user := model.LoginUser{Password: "333", Name: "ttt"}
	//tools.DB.Create(&user)
	//variable u model.LoginUser
	//tools.DB.Where("name=?", "ttt").First(&u)
	//Tlog.Println(u)
}
