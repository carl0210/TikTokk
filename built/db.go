package main

import (
	"TikTokk/model"
	"TikTokk/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	//utils.DB.AutoMigrate(&model.UserFollowed{})
	//utils.DB.AutoMigrate(&model.User{})
	//utils.DB.AutoMigrate(&model.UserFavorite{})
	//utils.DB.AutoMigrate(&model.Comment{})
	utils.DB.AutoMigrate(&model.Chat_Message{})
	//utils.DB.AutoMigrate(&model.Video{})

	// user := model.LoginUser{Password: "333", Name: "ttt"}
	//utils.DB.Create(&user)
	//var u model.LoginUser
	//utils.DB.Where("name=?", "ttt").First(&u)
	//log.Println(u)
}
