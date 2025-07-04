package model

import "bosh-admin/dao"

// SysMenu 系统菜单
type SysMenu struct {
    dao.BasicModel
    Path            string    `gorm:"not null;comment:'路由路径'" json:"path"`
    Name            string    `gorm:"type:varchar(30);not null;unique;comment:'路由名称'" json:"name"`
    Redirect        string    `gorm:"comment:'路由重定向'" json:"redirect"`
    Component       string    `gorm:"comment:'按需加载需要展示的页面'" json:"component"`
    ParentId        uint      `gorm:"default:0;comment:'父级菜单id'" json:"parentId"`
    MenuType        int       `gorm:"type:tinyint;default:0;comment:'菜单类型 0菜单 1iframe 2外链 3按钮'" json:"menuType"`
    Title           string    `gorm:"type:varchar(30);not null;comment:'菜单名称'" json:"title"`
    Icon            string    `gorm:"type:varchar(30);comment:'菜单图标'" json:"icon"`
    DisplayOrder    int       `gorm:"type:int unsigned;default:0;comment:'菜单排序'" json:"displayOrder"`
    ExtraIcon       string    `gorm:"type:varchar(30);comment:'菜单名称右侧的额外图标'" json:"extraIcon"`
    Transition      string    `gorm:"comment:'页面动画'" json:"transition"`
    EnterTransition string    `gorm:"comment:'当前页面进场动画'" json:"enterTransition"`
    LeaveTransition string    `gorm:"comment:'当前页面离场动画'" json:"leaveTransition"`
    ActivePath      string    `gorm:"comment:'菜单激活'" json:"activePath"`
    AuthCode        string    `gorm:"comment:'权限标识'" json:"authCode"`
    FrameSrc        string    `gorm:"comment:'需要内嵌的iframe链接地址'" json:"frameSrc"`
    FrameLoading    bool      `gorm:"type:tinyint(1);default:0;comment:'内嵌的iframe页面是否开启首次加载动画'" json:"frameLoading"`
    ShowLink        bool      `gorm:"type:tinyint(1);default:0;comment:'是否显示该菜单'" json:"showLink"`
    ShowParent      bool      `gorm:"type:tinyint(1);default:0;comment:'是否显示父级菜单'" json:"showParent"`
    KeepAlive       bool      `gorm:"type:tinyint(1);default:0;comment:'是否缓存改路由页面'" json:"keepAlive"`
    HiddenTag       bool      `gorm:"type:tinyint(1);default:0;comment:'当前菜单名称或自定义信息禁止添加到标签页'" json:"hiddenTag"`
    FixedTag        bool      `gorm:"type:tinyint(1);default:0;comment:'固定标签页'" json:"fixedTag"`
    Children        []SysMenu `gorm:"-" json:"children"`
}
