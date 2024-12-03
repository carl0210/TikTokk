package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"net/http"

	"github.com/gin-gonic/gin"
)

type t struct {
	Code string `json:"code" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Use(ValidateRequestMiddleware())

	router.POST("/hello", handleExampleRequest)

	router.Run(":8080")
}

func handleExampleRequest(c *gin.Context) {
	// 处理请求
	c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully"})
}

func ValidateRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在此处执行请求验证逻辑

		// 检查请求中是否包含非预期字段
		//var expectedFields = []string{"code"}
		//var requestData map[string]interface{}
		var ttt t
		if err := c.ShouldBindWith(&ttt, binding.FormPost); err != nil {
			fmt.Println("requestData=", ttt)
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			c.Abort()
			return
		}
		fmt.Println("requestData=", ttt)

		//for field := range requestData {
		//	if !contains(expectedFields, field) {
		//		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Unexpected field: %s", field)})
		//		c.Abort()
		//		return
		//	}
		//}

		c.Next()
	}
}

func contains(fields []string, target string) bool {
	for _, field := range fields {
		if field == target {
			return true
		}
	}
	return false
}
