package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"

    "github.com/gin-gonic/gin"
)

func SetBasicRouter(rg *gin.RouterGroup) {
    group := rg.Group("/basic")

    basic := api.NewBasicHandler()
    {
        group.GET("/health", ctx.Handler(basic.Health))
        group.GET("/captcha", ctx.Handler(basic.Captcha))
    }

    user := api.NewSysUserHandler()
    {
        group.POST("/login", ctx.Handler(user.Login))
        group.POST("/refreshToken", ctx.Handler(user.RefreshToken))
    }
}
