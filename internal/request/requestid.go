package request

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewRequestID 中间件生成请求ID
func NewRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get("RequestID"); exists {
		return requestID.(string)
	}
	// 如果没有找到 RequestID，返回一个空字符串或者日志记录一个警告
	return ""
}
