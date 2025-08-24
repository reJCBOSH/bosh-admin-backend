package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetDeptRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysDept")
    groupRecord := rg.Group("/sysDept").Use(middleware.OperationRecord())

    dept := handler.NewSysDeptHandler()
    {
        group.GET("/getTree", ctx.Handler(dept.GetDeptTree))
        group.GET("/getList", ctx.Handler(dept.GetDeptList))
        group.GET("/getInfo", ctx.Handler(dept.GetDeptInfo))

    }
    {
        groupRecord.POST("/add", ctx.Handler(dept.AddDept))
        groupRecord.POST("/edit", ctx.Handler(dept.EditDept))
        groupRecord.POST("/del", ctx.Handler(dept.DelDept))
    }
}
