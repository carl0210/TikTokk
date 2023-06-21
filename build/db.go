package main

import (
	"TikTokk/internal/TikTokk"
)

func main() {
	TikTokk.Config()
	TikTokk.Mysql()
	//tools.DB.AutoMigrate(&model.UserFollowed{})
	//tools.DB.AutoMigrate(&model.User{})
	//tools.DB.AutoMigrate(&model.UserFavorite{})
	//tools.DB.AutoMigrate(&model.Comment{})
	//TikTokk.DB.AutoMigrate(&model.Chat_Message{})
	//tools.DB.AutoMigrate(&model.Video{})

	// user := model.LoginUser{Password: "333", Name: "ttt"}
	//tools.DB.Create(&user)
	//variable u model.LoginUser
	//tools.DB.Where("name=?", "ttt").First(&u)
	//log.Println(u)
}
