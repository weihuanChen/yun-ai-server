package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"yinglian.com/yun-ai-server/configs"
	"yinglian.com/yun-ai-server/internal/db"
	"yinglian.com/yun-ai-server/pkg/router"
)

func Start() {
	runApp()

}

func runApp() {
	fmt.Println("Hello , AI App")
	r := router.AiServerRouter()
	db.New()
	// 自动迁移模型
	db.AutoMigrate()
	runHttpServer(r)
}
func runHttpServer(r *gin.Engine) {
	v := &http.Server{Handler: r}
	listener, err := net.Listen("tcp", ":"+configs.Cfg.App.Port)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个 channel 用于通知服务器启动成功
	started := make(chan struct{})

	go func() {
		fmt.Println("正在努力启动中...")
		// 通知服务
		close(started)
		if err := v.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("服务启动失败: %v", err)
		}
	}()
	// 等待服务器启动
	<-started
	fmt.Println("服务器成功启动在端口8099")

	// 服务退出 中断信号传入
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("收到关闭应用信号，正在下线服务器...")
	// 创建带超时的上下文 获取cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 确保cancel执行并释放资源
	defer cancel()

	// 关闭服务器
	if err := v.Shutdown(ctx); err != nil {
		log.Fatalf("服务下线失败: %v", err)
	}
	fmt.Println("服务已优雅停止...")
}
