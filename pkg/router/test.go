package router

import (
	"github.com/gin-gonic/gin"
	"yinglian.com/yun-ai-server/internal/utils"
)

func testRouter(r *gin.RouterGroup) {
	r.GET("/print", func(ctx *gin.Context) {
		utils.BizLogger(ctx).Infof("测试日志")
		ctx.JSON(200, gin.H{
			"message": "Test Router Group",
		})
	})
}
