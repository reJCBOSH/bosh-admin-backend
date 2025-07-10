package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetUserRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysUser")
    groupRecord := rg.Group("/sysUser").Use(middleware.OperationRecord())

    user := api.NewSysUserHandler()
    {
        group.GET("/getUserList", ctx.Handler(user.GetUserList))
        group.GET("/getUserInfo", ctx.Handler(user.GetUserInfo))
        group.GET("/getSelfInfo", ctx.Handler(user.GetSelfInfo))
    }
    {
        groupRecord.POST("/addUser", ctx.Handler(user.AddUser))
        groupRecord.POST("/editUser", ctx.Handler(user.EditUser))
        groupRecord.POST("/delUser", ctx.Handler(user.DelUser))
        groupRecord.POST("/resetPassword", ctx.Handler(user.ResetPassword))
        groupRecord.POST("/setUserStatus", ctx.Handler(user.SetUserStatus))
        groupRecord.POST("/editSelfInfo", ctx.Handler(user.EditSelfInfo))
        groupRecord.POST("/editSelfPassword", ctx.Handler(user.EditSelfPassword))
    }
}
