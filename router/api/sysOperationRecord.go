package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetOperationRecordRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysOperationRecord")
    groupRecord := rg.Group("/sysOperationRecord").Use(middleware.OperationRecord())

    operationRecord := api.NewSysOperationRecordHandler()
    {
        group.GET("/getList", ctx.Handler(operationRecord.GetOperationRecordList))
        group.GET("/getInfo", ctx.Handler(operationRecord.GetOperationRecordInfo))
    }
    {
        groupRecord.POST("/del", ctx.Handler(operationRecord.DelOperationRecord))
        groupRecord.POST("/batchDel", ctx.Handler(operationRecord.BatchDelOperationRecord))
    }
}
