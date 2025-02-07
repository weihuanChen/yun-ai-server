package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"slices"
	"strings"
	"time"
)

// 允许特定的来源
var corsAllows = []string{"线上doadmin"}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 允许的 HTTP 方法
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		// 允许的 HTTP 头部
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token"},
		// 暴露的 HTTP 头部
		ExposeHeaders: []string{"Content-Length"},
		// 是否允许携带身份凭证
		AllowCredentials: true,
		// 自定义允许的源
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 允许开发环境
				return true
			}
			// corsAllows 可以添加允许特定的来源
			return slices.Contains(corsAllows, origin)
		},
		// 最大缓存时间
		MaxAge: 12 * time.Hour,
	})
}
