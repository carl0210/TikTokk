package midware

import (
	"TikTokk/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthnByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := utils.ParseByQuery(c)
		if err != nil {
			log.Println("authn err=", err)
			c.JSON(404, gin.H{
				"status_code": 200,
				"status_msg":  "token失败",
			})
			c.Abort()
			return
		}
		log.Println("auth query value=", value)
		c.Set(utils.Config.IdentityKey, value)
		c.Next()

	}
}

func AuthnByBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := utils.ParseByBody(c)
		if err != nil {
			log.Println("authn err=", err)
			c.JSON(404, gin.H{
				"status_code": 200,
				"status_msg":  "token失败",
			})
			c.Abort()
			return
		}
		c.Set(utils.Config.IdentityKey, value)
		c.Next()

	}
}
