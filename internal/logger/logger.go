package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"sync"
	"time"
	"yinglian.com/yun-ai-server/configs"
	"yinglian.com/yun-ai-server/internal/global"
	"yinglian.com/yun-ai-server/internal/utils"
)

var once sync.Once // 只初始化一次
func initLogger() {
	// 确保日志只执行一次
	once.Do(func() {
		fmt.Println("初始化日志调用.....")
		logFilePath := configs.Cfg.Logger.LogFilePath
		logFileName := configs.Cfg.Logger.LogFileName
		logTimestampFmt := configs.Cfg.Logger.LogTimestampFmt
		logLevel := configs.Cfg.Logger.LogLevel
		logMaxAge := configs.Cfg.Logger.LogMaxAge
		logRotationTime := configs.Cfg.Logger.LogRotationTime

		// 创建日志目录
		if err := os.MkdirAll(logFilePath, 0755); err != nil {
			log.Fatalf("创建日志目录失败: %v", err)
		}
		// 初始化
		logger := logrus.New()
		// 设置日志格式
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: logTimestampFmt,
		})

		// 设置日志级别
		logLevelParsed, err := logrus.ParseLevel(logLevel)
		if err != nil {
			log.Fatalf("日志级别解析失败: %v", err)
		}
		logger.SetLevel(logLevelParsed)

		// 配置日志轮转
		fileName := path.Join(logFilePath, logFileName)
		maxAge := time.Duration(logMaxAge) * time.Hour
		rotationTime := time.Duration(logRotationTime) * time.Hour

		writer, err := rotatelogs.New(
			path.Join(logFilePath, "%Y%m%d.log"),
			rotatelogs.WithLinkName(fileName),
			rotatelogs.WithMaxAge(maxAge),
			rotatelogs.WithRotationTime(rotationTime),
		)
		if err != nil {
			log.Fatalf("设置日志轮转失败: %v", err)
		}

		// 配置日志级别与轮转日志的映射
		writerMap := lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.FatalLevel: writer,
			logrus.DebugLevel: writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.PanicLevel: writer,
		}

		// 添加 Hook，使用 JSONFormatter
		hook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{
			TimestampFormat: logTimestampFmt,
		})

		logger.SetOutput(os.Stdout) // 仅输出到控制台，文件输出由 hook 处理
		logger.AddHook(hook)

		global.SysLog = logger
	})
}

// 中间件，处理通用请求
func New() gin.HandlerFunc {
	initLogger() // 日志初始化

	return func(c *gin.Context) {
		reqId := c.GetString("RequestID")
		bizLog := global.SysLog.WithFields(
			logrus.Fields{
				"requestId": reqId,
				"requestIp": c.ClientIP(),
			})
		// 将 BizLog 存储到当前请求上下文中
		utils.SetSetBizLogger(c, bizLog)
		c.Next()
	}
}
