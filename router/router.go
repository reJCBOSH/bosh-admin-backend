package router

import (
	"bosh-admin/router/api"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	group := router.Group("/api")
	{
		api.SetBasicRouter(group)
		api.SetMenuRouter(group)
		api.SetDeptRouter(group)
		api.SetRoleRouter(group)
		api.SetUserRouter(group)
	}
}
