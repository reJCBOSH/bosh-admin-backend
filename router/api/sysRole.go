package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetRoleRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysRole")
    groupRecord := rg.Group("/sysRole").Use(middleware.OperationRecord())

    role := handler.NewSysRoleHandler()
    {
        group.GET("/getList", ctx.Handler(role.GetRoleList))
        group.GET("/getInfo", ctx.Handler(role.GetRoleInfo))
        group.GET("/getMenu", ctx.Handler(role.GetRoleMenu))
        group.GET("/getMenuIds", ctx.Handler(role.GetRoleMenuIds))
        group.GET("/getDeptIds", ctx.Handler(role.GetRoleDeptIds))
    }
    {
        groupRecord.POST("/add", ctx.Handler(role.AddRole))
        groupRecord.POST("/edit", ctx.Handler(role.EditRole))
        groupRecord.POST("/del", ctx.Handler(role.DelRole))
        groupRecord.POST("/setMenuAuth", ctx.Handler(role.SetRoleMenuAuth))
        groupRecord.POST("/setDataAuth", ctx.Handler(role.SetRoleDataAuth))
        groupRecord.POST("/setStatus", ctx.Handler(role.SetRoleStatus))
    }
}
