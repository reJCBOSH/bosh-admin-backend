package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetMenuRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysMenu")
    groupRecord := rg.Group("/sysMenu").Use(middleware.OperationRecord())

    menu := handler.NewSysMenuHandler()
    {
        group.GET("/getTree", ctx.Handler(menu.GetMenuTree))
        group.GET("/getList", ctx.Handler(menu.GetMenuList))
        group.GET("/getInfo", ctx.Handler(menu.GetMenuInfo))
        group.POST("/getAsyncRoutes", ctx.Handler(menu.GetAsyncRoutes))
    }
    {
        groupRecord.POST("/add", ctx.Handler(menu.AddMenu))
        groupRecord.POST("/edit", ctx.Handler(menu.EditMenu))
        groupRecord.POST("/del", ctx.Handler(menu.DelMenu))
    }
}
