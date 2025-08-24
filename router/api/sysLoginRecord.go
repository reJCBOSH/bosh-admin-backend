package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetLoginRecordRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysLoginRecord")
    groupRecord := rg.Group("/sysLoginRecord").Use(middleware.OperationRecord())

    loginRecord := handler.NewSysLoginRecordHandler()
    {
        group.GET("/getList", ctx.Handler(loginRecord.GetLoginRecordList))

    }
    {
        groupRecord.POST("/del", ctx.Handler(loginRecord.DelLoginRecord))
        groupRecord.POST("/batchDel", ctx.Handler(loginRecord.BatchDelLoginRecord))
    }
}
