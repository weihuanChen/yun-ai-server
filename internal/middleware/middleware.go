package middleware

import (
	"github.com/gin-gonic/gin"
	"yinglian.com/yun-ai-server/internal/logger"
	"yinglian.com/yun-ai-server/internal/request"
)

func InitMiddleware(r *gin.Engine) {
	r.Use(logger.New())
	// 全局 ID
	r.Use(request.NewRequestID())
	// cors 中间件
	r.Use(Cors())
}
