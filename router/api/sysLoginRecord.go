package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "github.com/gin-gonic/gin"
)

func SetLoginRecordRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysLoginRecord")

    loginRecord := api.NewSysLoginRecordHandler()
    {
        group.GET("/getLoginRecordList", ctx.Handler(loginRecord.GetLoginRecordList))
        group.POST("/delLoginRecord", ctx.Handler(loginRecord.DelLoginRecord))
        group.POST("/batchDel", ctx.Handler(loginRecord.BatchDelLoginRecord))
    }
}
