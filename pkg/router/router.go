package router

import (
	"github.com/gin-gonic/gin"
)

func AiServerRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// time.Sleep(10 * time.Second) // 模拟长时间处理的请求
		c.JSON(200, gin.H{
			"message": "Cool, AI App 🎉 !",
		})
	})
	return r
}
