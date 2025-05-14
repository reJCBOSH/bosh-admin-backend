package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/handler/api"
	"github.com/gin-gonic/gin"
)

func SetRoleRouter(r *gin.RouterGroup) {
	group := r.Group("/sysRole")

	role := api.NewSysRoleHandler()
	{
		group.GET("/getRoleList", ctx.Handler(role.GetRoleList))
		group.GET("/getRoleInfo", ctx.Handler(role.GetRoleInfo))
		group.POST("/addRole", ctx.Handler(role.AddRole))
		group.POST("/editRole", ctx.Handler(role.EditRole))
		group.POST("/delRole", ctx.Handler(role.DelRole))
		group.GET("/getRoleMenu", ctx.Handler(role.GetRoleMenu))
		group.GET("/getRoleMenuIds", ctx.Handler(role.GetRoleMenuIds))
		group.POST("/setRoleMenuAuth", ctx.Handler(role.SetRoleMenuAuth))
		group.GET("/getRoleDeptIds", ctx.Handler(role.GetRoleDeptIds))
		group.POST("/setRoleDataAuth", ctx.Handler(role.SetRoleDataAuth))
		group.POST("/setRoleStatus", ctx.Handler(role.SetRoleStatus))
	}
}
