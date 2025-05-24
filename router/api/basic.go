package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/handler/api"
	"github.com/gin-gonic/gin"
)

func SetBasicRouter(r *gin.RouterGroup) {
	group := r.Group("/basic")

	user := api.NewSysUserHandler()
	{
		group.POST("/login", ctx.Handler(user.Login))
		group.POST("/refreshToken", ctx.Handler(user.RefreshToken))
	}
}
