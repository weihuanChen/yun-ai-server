package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"yinglian.com/yun-ai-server/internal/global"
)

const (
	bizLogKey = "BizLog"
)

// 返回一个有效的日志对象，避免出现空指针
func BizLogger(c *gin.Context) *logrus.Entry {
	// 从gin上下文中获取BizLog
	if entry, exists := c.Get(bizLogKey); exists {
		// 如果存在并且非空返回
		if logEntry, ok := entry.(*logrus.Entry); ok && logEntry != nil {
			return logEntry
		}
	}
	// 没有找到就返回全局日志
	return logrus.NewEntry(global.SysLog)
}

// 设置到gin的上下文
func SetSetBizLogger(c *gin.Context, logger *logrus.Entry) {
	c.Set(bizLogKey, logger)
}
