package router

import (
	"github.com/gin-gonic/gin"
	"time"
)

func AiServerRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		// time.Sleep(10 * time.Second) // 模拟长时间处理的请求
		c.JSON(200, gin.H{
			"message": "Cool, AI App 🎉 !",
		})
	})
	r.GET("/long", func(c *gin.Context) {
		time.Sleep(10 * time.Second) // 模拟长时间处理的请求
		c.JSON(200, gin.H{
			"message": "Cool, AI App 🎉 !",
		})
	})
	// 统一前缀
	v1 := r.Group("/api/v1")

	// 测试路由
	testGroup := v1.Group("/testGroup")
	testRouter(testGroup)
	return r
}
