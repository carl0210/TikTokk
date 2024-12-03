package token

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type C struct {
	Key         string
	IdentityKey string
}

var AuthMissError = fmt.Errorf("鉴权方法不同")

var Config = C{"", ""}

func Parse(tokenString string, key string) (value string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", AuthMissError
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		value = claims[Config.IdentityKey].(string)
	} else {
		return "", err
	}

	return value, nil
}

func ParseByQuery(c *gin.Context) (value string, err error) {
	//a := c.Request.Header.GetByAuthorID("token")
	a := c.Query("token")
	if len(a) == 0 {
		return "", AuthMissError
	}
	return Parse(a, Config.Key)
}

func ParseByBody(c *gin.Context) (value string, err error) {
	//a := c.Request.Header.GetByAuthorID("token")
	a := c.PostForm("token")
	if len(a) == 0 {
		return "", AuthMissError
	}
	return Parse(a, Config.Key)
}

func Sign(identity string) string {
	c := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		Config.IdentityKey: identity,
		"exp":              time.Now().Add(12 * time.Hour).Unix(),
		"iat":              time.Now().Unix(),
		"nbf":              time.Now().Unix(),
	})
	s, err := c.SignedString([]byte(Config.Key))
	if err != nil {
		return ""
	}
	return s
}
