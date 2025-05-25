package model

import "bosh-admin/dao"

// SysDept 部门
type SysDept struct {
	dao.BasicModel
	ParentId     uint      `gorm:"parent_id" json:"parentId"`         // 父级id
	DeptName     string    `gorm:"dept_name" json:"deptName"`         // 部门名称
	DeptCode     string    `gorm:"dept_code" json:"deptCode"`         // 部门标识
	DeptPath     string    `gorm:"dept_path" json:"deptPath"`         // 部门路径
	Remark       string    `gorm:"remark" json:"remark"`              // 备注
	Status       int       `gorm:"status" json:"status"`              // 状态 0停用 1启用
	DisplayOrder int       `gorm:"display_order" json:"displayOrder"` // 显示顺序
	Children     []SysDept `gorm:"-" json:"children"`
}
