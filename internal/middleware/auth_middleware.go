package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/utils"
)

func NewAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userId *int64 = nil
		// 从请求handler中去token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 设置默认值
			defaultUserId := int64(0)
			userId = &defaultUserId
			// 尝试提取用户信息
			tokenStr := strings.TrimPrefix(authHeader, global.BearerPrefix)
			if claims, err := utils.ValidateToken(tokenStr); err == nil {
				// 验证成功 获取userid
				potentialUserId := claims.UserId
				// 检验缓存是否生效
				accKey := global.AccAuthTokenCachePrefix + strconv.FormatInt(potentialUserId, 10)
				conn := global.RDB.Get()
				defer conn.Close()
				if accTokenValue, err := conn.Do(global.GetCmd, accKey); err == nil {
					// Redis查询成功
					if accTokenValueStr, err := redis.String(accTokenValue, nil); err == nil {
						// Token 匹配
						if accTokenValueStr != "" && accTokenValueStr == tokenStr {
							// 所有验证都通过，设置实际的 userId
							userId = &potentialUserId
						}
					}
				}
			}
		}
		// 设置 userId（传递指针）到上下文并继续请求流程
		c.Set(global.LocalsUserIdKey, userId)
		c.Next()
	}
}
