package main

import (
	"TikTokk/route"
	"TikTokk/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitJwt(viper.GetString("jwt.key"), viper.GetString("jwt.identityKey"))
	utils.InitStore()
	utils.InitFeedlen()
	utils.InitFileSavePath()

	engine := route.Route()
	engine.Run(":6666")
}
