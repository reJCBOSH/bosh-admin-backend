package dto

import (
	"bosh-admin/dao"
	"bosh-admin/dao/model"
)

type GetDeptListRequest struct {
	Pagination
	DeptName string `json:"deptName" form:"deptName"`
	DeptCode string `json:"deptCode" form:"deptCode"`
}

type GetDeptListResponse struct {
	List  []model.SysDept `json:"list"`
	Total int64           `json:"total"`
}

type AddDeptRequest struct {
	dao.AddBasicModel
	ParentId     uint   `json:"parentId" form:"parentId" binding:"gte=0"`                  // 父级id
	DeptName     string `json:"deptName" form:"deptName" binding:"required"`               // 部门名称
	DeptCode     string `json:"deptCode" form:"deptCode" binding:"required"`               // 部门标识
	DeptPath     string `json:"deptPath" form:"deptPath" binding:"required"`               // 部门路径
	Remark       string `json:"remark" form:"remark"`                                      // 备注
	Status       int    `json:"status" form:"status" binding:"oneof=0 1"`                  // 状态 0停用 1启用
	DisplayOrder int    `json:"displayOrder" form:"displayOrder" binding:"gte=0,lte=9999"` // 显示顺序
}

type EditDeptRequest struct {
	dao.EditBasicModel
	DeptName     string `json:"deptName" form:"deptName" binding:"required"`               // 部门名称
	Remark       string `json:"remark" form:"remark"`                                      // 备注
	Status       int    `json:"status" form:"status" binding:"oneof=0 1"`                  // 状态 0停用 1启用
	DisplayOrder int    `json:"displayOrder" form:"displayOrder" binding:"gte=0,lte=9999"` // 显示顺序
}
