package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetRoleRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysRole")
    groupRecord := rg.Group("/sysRole").Use(middleware.OperationRecord())

    role := api.NewSysRoleHandler()
    {
        group.GET("/getRoleList", ctx.Handler(role.GetRoleList))
        group.GET("/getRoleInfo", ctx.Handler(role.GetRoleInfo))
        group.GET("/getRoleMenu", ctx.Handler(role.GetRoleMenu))
        group.GET("/getRoleMenuIds", ctx.Handler(role.GetRoleMenuIds))
        group.GET("/getRoleDeptIds", ctx.Handler(role.GetRoleDeptIds))
    }
    {
        groupRecord.POST("/addRole", ctx.Handler(role.AddRole))
        groupRecord.POST("/editRole", ctx.Handler(role.EditRole))
        groupRecord.POST("/delRole", ctx.Handler(role.DelRole))
        groupRecord.POST("/setRoleMenuAuth", ctx.Handler(role.SetRoleMenuAuth))
        groupRecord.POST("/setRoleDataAuth", ctx.Handler(role.SetRoleDataAuth))
        groupRecord.POST("/setRoleStatus", ctx.Handler(role.SetRoleStatus))
    }
}
