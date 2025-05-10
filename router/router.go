package router

import (
	"bosh-admin/router/api"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	group := router.Group("/api")
	{
		api.SetMenuRouter(group)
	}
}
