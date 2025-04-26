package main

import (
	"errors"
	"fmt"
	"net/http"

	"bosh-admin/core/log"
	"bosh-admin/global"
	"bosh-admin/initialize"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	initialize.InitConfig()
	// 初始化日志
	initialize.InitLog()

	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

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
