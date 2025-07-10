package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "bosh-admin/middleware"

    "github.com/gin-gonic/gin"
)

func SetDeptRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysDept")
    groupRecord := rg.Group("/sysDept").Use(middleware.OperationRecord())

    dept := api.NewSysDeptHandler()
    {
        group.GET("/getDeptTree", ctx.Handler(dept.GetDeptTree))
        group.GET("/getDeptList", ctx.Handler(dept.GetDeptList))
        group.GET("/getDeptInfo", ctx.Handler(dept.GetDeptInfo))

    }
    {
        groupRecord.POST("/addDept", ctx.Handler(dept.AddDept))
        groupRecord.POST("/editDept", ctx.Handler(dept.EditDept))
        groupRecord.POST("/delDept", ctx.Handler(dept.DelDept))
    }
}
