package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/handler/api"

	"github.com/gin-gonic/gin"
)

func SetMenuRouter(r *gin.RouterGroup) {
	group := r.Group("/sysMenu")

	menu := api.NewSysMenuHandler()
	{
		group.GET("/getMenuTree", ctx.Handler(menu.GetMenuTree))
		group.GET("/getMenuList", ctx.Handler(menu.GetMenuList))
		group.GET("/getMenuInfo", ctx.Handler(menu.GetMenuInfo))
	}
}
