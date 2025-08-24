package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetUserRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysUser")
    groupRecord := rg.Group("/sysUser").Use(middleware.OperationRecord())

    user := handler.NewSysUserHandler()
    {
        group.GET("/getList", ctx.Handler(user.GetUserList))
        group.GET("/getInfo", ctx.Handler(user.GetUserInfo))
        group.GET("/getSelfInfo", ctx.Handler(user.GetSelfInfo))
    }
    {
        groupRecord.POST("/add", ctx.Handler(user.AddUser))
        groupRecord.POST("/edit", ctx.Handler(user.EditUser))
        groupRecord.POST("/del", ctx.Handler(user.DelUser))
        groupRecord.POST("/resetPassword", ctx.Handler(user.ResetPassword))
        groupRecord.POST("/setStatus", ctx.Handler(user.SetUserStatus))
        groupRecord.POST("/editSelfInfo", ctx.Handler(user.EditSelfInfo))
        groupRecord.POST("/editSelfPassword", ctx.Handler(user.EditSelfPassword))
    }
}
