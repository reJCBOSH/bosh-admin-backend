package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"

    "github.com/gin-gonic/gin"
)

func SetMenuRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysMenu")

    menu := api.NewSysMenuHandler()
    {
        group.GET("/getMenuTree", ctx.Handler(menu.GetMenuTree))
        group.GET("/getMenuList", ctx.Handler(menu.GetMenuList))
        group.GET("/getMenuInfo", ctx.Handler(menu.GetMenuInfo))
        group.POST("/getAsyncRoutes", ctx.Handler(menu.GetAsyncRoutes))
        group.POST("/addMenu", ctx.Handler(menu.AddMenu))
        group.POST("/editMenu", ctx.Handler(menu.EditMenu))
        group.POST("/delMenu", ctx.Handler(menu.DelMenu))
    }
}
