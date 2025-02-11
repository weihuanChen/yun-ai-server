package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	biz_err "yinglian.com/yun-ai-server/internal/error"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/utils"
	"yinglian.com/yun-ai-server/pkg/vo"
)

func testRouter(r *gin.RouterGroup) {
	r.GET("/print", func(ctx *gin.Context) {
		utils.BizLogger(ctx).Infof("测试日志")
		ctx.JSON(200, gin.H{
			"message": "Test Router Group",
		})
	})
	r.GET("/testRedis", func(ctx *gin.Context) {
		utils.BizLogger(ctx).Infof("开始写入缓存...")
		// 初始化缓存连接
		conn := global.RDB.Get()
		defer conn.Close()
		_, err := conn.Do("SET", "TEST:", "测试 value")
		if err != nil {
			utils.BizLogger(ctx).Errorf("测试写入缓存失败: %v", err)
			ctx.JSON(500, gin.H{
				"message": "测试写入缓存失败",
			})
		}
		utils.BizLogger(ctx).Infof("写入缓存成功...")

		utils.BizLogger(ctx).Infof("开始读取缓存...")
		// 这里可以复用 conn 打开的连接
		testCache, err := conn.Do("GET", "TEST:")
		if err != nil {
			utils.BizLogger(ctx).Errorf("测试读取缓存失败: %v", err)
			ctx.JSON(500, gin.H{
				"message": "测试读取缓存失败",
			})
		}
		utils.BizLogger(ctx).Infof("读取缓存成功, key: %s , value: %s", "TEST:", testCache)
		str, _ := redis.String(testCache, nil)
		ctx.JSON(200, gin.H{
			"message": str,
		})
	})
	r.GET("/testVOSuccess", func(ctx *gin.Context) {
		ctx.JSON(200, vo.Success(nil, "", ctx))
	})

	r.GET("/testVOFail", func(ctx *gin.Context) {
		ctx.JSON(200, vo.Fail(biz_err.New(biz_err.ServerError), nil, ctx))
	})

}
