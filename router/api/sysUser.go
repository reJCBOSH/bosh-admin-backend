package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/handler/api"
	"github.com/gin-gonic/gin"
)

func SetUserRouter(r *gin.RouterGroup) {
	group := r.Group("/sysUser")

	user := api.NewSysUserHandler()
	{
		group.GET("/getUserList", ctx.Handler(user.GetUserList))
		group.GET("/getUserInfo", ctx.Handler(user.GetUserInfo))
		group.POST("/addUser", ctx.Handler(user.AddUser))
		group.POST("/editUser", ctx.Handler(user.EditUser))
		group.POST("/delUser", ctx.Handler(user.DelUser))
	}
}
