package router

import (
    "bosh-admin/middleware"
    "bosh-admin/router/api"

    "github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
    group := router.Group("/api")

    public := group.Group("")
    {
        api.SetBasicRouter(public)
    }

    private := group.Group("")
    private.Use(middleware.JWTApiAuth())
    {
        api.SetMenuRouter(private)
        api.SetDeptRouter(private)
        api.SetRoleRouter(private)
        api.SetUserRouter(private)
        api.SetLoginRecordRouter(private)
        api.SetOperationRecordRouter(private)
    }
}
