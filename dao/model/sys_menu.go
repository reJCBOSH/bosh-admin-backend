package model

import "bosh-admin/dao"

// SysMenu 系统菜单
type SysMenu struct {
	dao.BasicModel
	Path            string    `gorm:"path" json:"path"`                        // 路由路径
	Name            string    `gorm:"name" json:"name"`                        // 路由名称(必须保持唯一)
	Redirect        string    `gorm:"redirect" json:"redirect"`                // 路由重定向
	Component       string    `gorm:"component" json:"component"`              // 按需加载需要展示的页面
	ParentId        uint      `gorm:"parent_id" json:"parentId"`               // 父级菜单id
	MenuType        int       `gorm:"menu_type" json:"menuType"`               // 菜单类型 0菜单 1iframe 2外链 3按钮
	Title           string    `gorm:"title" json:"title"`                      // 菜单名称
	Icon            string    `gorm:"icon" json:"icon"`                        // 菜单图标
	DisplayOrder    int       `gorm:"display_order" json:"displayOrder"`       // 菜单排序
	ExtraIcon       string    `gorm:"extra_icon" json:"extraIcon"`             // 菜单名称右侧的额外图标
	Transition      string    `gorm:"transition" json:"transition"`            // 页面动画
	EnterTransition string    `gorm:"enter_transition" json:"enterTransition"` // 当前页面进场动画
	LeaveTransition string    `gorm:"leave_transition" json:"leaveTransition"` // 当前页面离场动画
	ActivePath      string    `gorm:"active_path" json:"activePath"`           // 菜单激活
	AuthCode        string    `gorm:"auth_code" json:"authCode"`               // 权限标识
	FrameSrc        string    `gorm:"frame_src" json:"frameSrc"`               // 需要内嵌的iframe链接地址
	FrameLoading    bool      `gorm:"frame_loading" json:"frameLoading"`       // 内嵌的iframe页面是否开启首次加载动画
	ShowLink        bool      `gorm:"show_link" json:"showLink"`               // 是否显示该菜单
	ShowParent      bool      `gorm:"show_parent" json:"showParent"`           // 是否显示父级菜单
	KeepAlive       bool      `gorm:"keep_alive" json:"keepAlive"`             // 是否缓存改路由页面
	HiddenTag       bool      `gorm:"hidden_tag" json:"hiddenTag"`             // 当前菜单名称或自定义信息禁止添加到标签页
	FixedTag        bool      `gorm:"fixed_tag" json:"fixedTag"`               // 固定标签页
	Children        []SysMenu `gorm:"-" json:"children"`                       // 子菜单
}

// TableName sys_menu
func (SysMenu) TableName() string {
	return "sys_menu"
}
