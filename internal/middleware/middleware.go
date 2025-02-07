package middleware

import (
	"github.com/gin-gonic/gin"
	"yinglian.com/yun-ai-server/internal/logger"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(logger.New())
	// cors 中间件
	r.Use(Cors())
}
