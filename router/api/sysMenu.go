package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetMenuRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysMenu")
    groupRecord := rg.Group("/sysMenu").Use(middleware.OperationRecord())

    menu := api.NewSysMenuHandler()
    {
        group.GET("/getMenuTree", ctx.Handler(menu.GetMenuTree))
        group.GET("/getMenuList", ctx.Handler(menu.GetMenuList))
        group.GET("/getMenuInfo", ctx.Handler(menu.GetMenuInfo))
        group.POST("/getAsyncRoutes", ctx.Handler(menu.GetAsyncRoutes))
    }
    {
        groupRecord.POST("/addMenu", ctx.Handler(menu.AddMenu))
        groupRecord.POST("/editMenu", ctx.Handler(menu.EditMenu))
        groupRecord.POST("/delMenu", ctx.Handler(menu.DelMenu))
    }
}
