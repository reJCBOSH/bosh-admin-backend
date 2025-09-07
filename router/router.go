package router

import (
	"bosh-admin/core/ctx"
	"bosh-admin/global"
	"bosh-admin/middleware"
	"bosh-admin/router/api"
	
	"github.com/gin-gonic/gin"
)

func SetStaticRouter(router *gin.Engine) {
	router.StaticFS("/"+global.Config.Local.Path, gin.Dir(global.Config.Local.StorePath, false))
}

func SetWebSocketRouter(router *gin.Engine) {
	router.GET("/ws", ctx.Handler(func(c *ctx.Context) {
		username := c.Query("username")
		global.WsHub.HandleConnection(c.Writer, c.Request, username)
	}))
}

func SetApiRouter(router *gin.Engine) {
	group := router.Group("/api")

	public := group.Group("")
	{
		api.SetBasicRouter(public)
	}

	private := group.Group("")
	private.Use(middleware.JWTApiAuth())
	{
		api.SetMenuRouter(private)
		api.SetDeptRouter(private)
		api.SetRoleRouter(private)
		api.SetUserRouter(private)
		api.SetLoginRecordRouter(private)
		api.SetOperationRecordRouter(private)
		api.SetApiRouter(private)
	}
}
