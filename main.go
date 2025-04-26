package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/initialize"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLog()
	// 初始化路由
	router := initialize.InitRouter()

	addr := fmt.Sprintf(":%d", global.Config.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	time.Sleep(10 * time.Microsecond)
	log.Info("服务已启动", addr)

	// 优雅关机
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("监听: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Info("关闭服务中...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("服务关闭失败: ", err)
	}

	log.Info("服务退出")
}
