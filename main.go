package main

import (
	"TikTokk/route"
	"TikTokk/utils"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitJwt(viper.GetString("jwt.key"), viper.GetString("jwt.identityKey"))
	utils.InitStore()
	utils.InitFeedlen()
	utils.InitFileSavePath()
	fmt.Println("123")
	engine := route.Route()
	engine.Run(":6666")
}
