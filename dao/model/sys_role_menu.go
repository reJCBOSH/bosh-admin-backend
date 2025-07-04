package model

// SysRoleMenu 角色-菜单关联
type SysRoleMenu struct {
    RoleId uint `gorm:"comment:角色id"`
    MenuId uint `gorm:"comment:菜单id"`
}
