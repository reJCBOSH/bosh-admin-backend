package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetOperationRecord(rg *gin.RouterGroup) {
    group := rg.Group("/sysOperationRecord")
    groupRecord := rg.Group("/sysOperationRecordRecord").Use(middleware.OperationRecord())

    operationRecord := api.NewSysOperationRecordHandler()
    {
        group.GET("/getOperationRecordList", ctx.Handler(operationRecord.GetOperationRecordList))
        group.GET("/getOperationRecordInfo", ctx.Handler(operationRecord.GetOperationRecordInfo))
    }
    {
        groupRecord.POST("/delOperationRecord", ctx.Handler(operationRecord.DelOperationRecord))
        groupRecord.POST("/batchDelOperationRecord", ctx.Handler(operationRecord.BatchDelOperationRecord))
    }
}
