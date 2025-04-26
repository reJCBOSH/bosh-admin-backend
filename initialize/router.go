package initialize

import (
	"net/http"

	"bosh-admin/core/log"
	"bosh-admin/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	if utils.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 使用gin默认Logger、Recovery中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 健康检测
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	log.Info("路由注册完成")
	return r
}
