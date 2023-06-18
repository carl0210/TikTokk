package main

import (
	"TikTokk/route"
	"TikTokk/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitJwt()
	utils.InitStore()
	utils.InitFeedlen()
	utils.InitSavePath()
	engine := route.Route()
	engine.Run(":6666")
}
