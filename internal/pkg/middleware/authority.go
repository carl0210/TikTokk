package middleware

import (
	"TikTokk/internal/pkg/Tlog"
	"TikTokk/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthnByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := token.ParseByQuery(c)
		if err != nil {
			c.JSON(404, gin.H{
				"status_code": 200,
				"status_msg":  "token失败",
			})
			c.Abort()
			return
		}
		c.Set(token.Config.IdentityKey, value)
		c.Next()

	}
}

func AuthnByBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := token.ParseByBody(c)
		if err != nil {
			Tlog.Infow(err.Error())
			c.JSON(404, gin.H{
				"status_code": 200,
				"status_msg":  "token失败",
			})
			c.Abort()
			return
		}
		c.Set(token.Config.IdentityKey, value)
		c.Next()

	}
}
