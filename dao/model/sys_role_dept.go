package model

// SysRoleDept 角色-部门关联
type SysRoleDept struct {
    RoleId uint `gorm:"comment:角色id"`
    DeptId uint `gorm:"comment:部门id"`
}
