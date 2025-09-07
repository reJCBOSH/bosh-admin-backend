package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/handler"
	"bosh-admin/middleware"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(rg *gin.RouterGroup) {
	group := rg.Group("/sysApi")
	groupRecord := rg.Group("/sysApi").Use(middleware.OperationRecord())

	sysApi := handler.NewSysApiHandler()
	{
		group.GET("/getApiList", ctx.Handler(sysApi.GetApiList))
		group.GET("/getApiGroups", ctx.Handler(sysApi.GetApiGroups))
	}
	{
		groupRecord.POST("/addApi", ctx.Handler(sysApi.AddApi))
		groupRecord.POST("/editApi", ctx.Handler(sysApi.EditApi))
		groupRecord.POST("/delApi", ctx.Handler(sysApi.DelApi))
	}
}
