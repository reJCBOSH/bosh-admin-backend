package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/handler/api"
    "github.com/gin-gonic/gin"
)

func SetDeptRouter(rg *gin.RouterGroup) {
    group := rg.Group("/sysDept")

    dept := api.NewSysDeptHandler()
    {
        group.GET("/getDeptTree", ctx.Handler(dept.GetDeptTree))
        group.GET("/getDeptList", ctx.Handler(dept.GetDeptList))
        group.GET("/getDeptInfo", ctx.Handler(dept.GetDeptInfo))
        group.POST("/addDept", ctx.Handler(dept.AddDept))
        group.POST("/editDept", ctx.Handler(dept.EditDept))
        group.POST("/delDept", ctx.Handler(dept.DelDept))
    }
}
