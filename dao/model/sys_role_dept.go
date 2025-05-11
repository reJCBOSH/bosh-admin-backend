package model

// SysRoleDept 角色-部门关联
type SysRoleDept struct {
	RoleId uint `gorm:"role_id"` // 角色Id
	DeptId uint `gorm:"dept_id"` // 部门Id
}

// TableName sys_role_dept
func (SysRoleDept) TableName() string {
	return "sys_role_dept"
}
