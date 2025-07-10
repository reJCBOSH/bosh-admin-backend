package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "github.com/gin-gonic/gin"
)

func SetOperationRecord(rg *gin.RouterGroup) {
    group := rg.Group("/sysOperationRecord")

    operationRecord := api.NewSysOperationRecordHandler()
    {
        group.GET("/getOperationRecordList", ctx.Handler(operationRecord.GetOperationRecordList))
        group.GET("/getOperationRecordInfo", ctx.Handler(operationRecord.GetOperationRecordInfo))
        group.POST("/delOperationRecord", ctx.Handler(operationRecord.DelOperationRecord))
        group.POST("/batchDelOperationRecord", ctx.Handler(operationRecord.BatchDelOperationRecord))
    }
}
