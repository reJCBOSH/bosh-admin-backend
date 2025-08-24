package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/handler"
	"bosh-admin/middleware"
	"github.com/gin-gonic/gin"
)

func SetBasicRouter(rg *gin.RouterGroup) {
	group := rg.Group("/basic")

	basic := handler.NewBasicHandler()
	{
		group.GET("/health", middleware.RateLimiter(5, 10), ctx.Handler(basic.Health))
		group.GET("/captcha", ctx.Handler(basic.Captcha))
		group.POST("/upload", ctx.Handler(basic.Upload))
	}

	user := handler.NewSysUserHandler()
	{
		group.POST("/login", ctx.Handler(user.Login))
		group.POST("/refreshToken", ctx.Handler(user.RefreshToken))
	}
}
