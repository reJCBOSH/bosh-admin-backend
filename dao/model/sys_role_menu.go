package model

// SysRoleMenu 角色-菜单关联
type SysRoleMenu struct {
	RoleId uint `gorm:"role_id"` // 角色id
	MenuId uint `gorm:"menu_id"` // 菜单id
}
