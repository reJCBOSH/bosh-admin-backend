package main

import (
	"errors"
	"fmt"
	"net/http"

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
	r := initialize.InitRouter()

	addr := fmt.Sprintf(":%d", global.Config.Server.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Info("服务已启动", addr)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err.Error())
	}
}
