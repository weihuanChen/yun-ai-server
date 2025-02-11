package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/middleware"
	"yinglian.com/yun-ai-server/pkg/serve/wire"
)

func AiServerRouter() *gin.Engine {
	r := gin.Default()
	middleware.InitMiddleware(r)
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
	initCtl(v1)
	return r
}
func initCtl(rg *gin.RouterGroup) {
	initAccountCtl(rg)
}
func initAccountCtl(rg *gin.RouterGroup) {
	accCtl, err := wire.InitializeAccountController(global.DB)
	if err != nil {
		log.Fatalf("初始化 AccountCtl 失败: %v", err)
	}
	accGroup := rg.Group("/account")
	accountRouter(accCtl, accGroup)
}
