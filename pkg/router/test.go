package router

import "github.com/gin-gonic/gin"

func testRouter(r *gin.RouterGroup) {
	r.GET("/print", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Test Router Group",
		})
	})
}
