package dto

import "bosh-admin/dao"

type GetMenuListReq struct {
    Pagination
    Title string `json:"title" form:"title"`
}

type MenuBasicItem struct {
    Path            string `json:"path" form:"path" validate:"required_if=MenuType 0|required_if=MenuType 1|required_if=MenuType 2"` // 路由路径
    Name            string `json:"name" form:"name" validate:"required_if=MenuType 0|required_if=MenuType 1|required_if=MenuType 2"` // 路由名称(必须保持唯一)
    Redirect        string `json:"redirect" form:"redirect"`                                                                         // 路由重定向
    Component       string `json:"component" form:"component"`                                                                       // 按需加载需要展示的页面
    ParentId        uint   `json:"parentId" form:"parentId"`                                                                         // 父级菜单id
    MenuType        int    `json:"menuType" form:"menuType" validate:"oneof=0 1 2 3"`                                                // 菜单类型 0菜单 1iframe 2外链 3按钮
    Title           string `json:"title" form:"title" validate:"required"`                                                           // 菜单名称
    Icon            string `json:"icon" form:"icon"`                                                                                 // 菜单图标
    DisplayOrder    int    `json:"displayOrder" form:"displayOrder"`                                                                 // 菜单排序
    ExtraIcon       string `json:"extraIcon" form:"extraIcon"`                                                                       // 菜单名称右侧的额外图标
    Transition      string `json:"transition" form:"transition"`                                                                     // 页面动画
    EnterTransition string `json:"enterTransition" form:"enterTransition"`                                                           // 当前页面进场动画
    LeaveTransition string `json:"leaveTransition" form:"leaveTransition"`                                                           // 当前页面离场动画
    ActivePath      string `json:"activePath" form:"activePath"`                                                                     // 菜单激活
    AuthCode        string `json:"authCode" form:"authCode" validate:"required_if=MenuType 3"`                                       // 权限标识
    FrameSrc        string `json:"frameSrc" form:"frameSrc"`                                                                         // 需要内嵌的iframe链接地址
    FrameLoading    bool   `json:"frameLoading" form:"frameLoading" `                                                                // 内嵌的iframe页面是否开启首次加载动画
    ShowLink        bool   `json:"showLink" form:"showLink"`                                                                         // 是否显示该菜单
    ShowParent      bool   `json:"showParent" form:"showParent"`                                                                     // 是否显示父级菜单
    KeepAlive       bool   `json:"keepAlive" form:"keepAlive"`                                                                       // 是否缓存改路由页面
    HiddenTag       bool   `json:"hiddenTag" form:"hiddenTag"`                                                                       // 当前菜单名称或自定义信息禁止添加到标签页
    FixedTag        bool   `json:"fixedTag" form:"fixedTag"`                                                                         // 固定标签页
}

type MenuItem struct {
    dao.BasicModel
    MenuBasicItem
}

type GetMenuDetailResp struct {
    MenuItem
}

type AddMenuReq struct {
    dao.AddBasicModel
    MenuBasicItem
}

type EditMenuReq struct {
    dao.EditBasicModel
    MenuBasicItem
}

// PureMenu Pure Admin 路由菜单
type PureMenu struct {
    Id        uint         `json:"-"`                   // 菜单Id
    ParentId  uint         `json:"-"`                   // 父级菜单Id
    Path      string       `json:"path"`                // 路由地址
    Name      string       `json:"name"`                // 路由名称(必须保持唯一)
    Redirect  string       `json:"redirect,omitempty"`  // 路由重定向
    Component string       `json:"component,omitempty"` // 按需加载需要展示的页面
    Meta      PureMenuMeta `json:"meta"`                // 路由元信息
    Children  []PureMenu   `json:"children,omitempty"`  // 子路由配置项
}

// PureMenuMeta Pure Admin 路由元信息
type PureMenuMeta struct {
    Title        string             `json:"title"`                  // 菜单名称
    Icon         string             `json:"icon"`                   // 菜单图标
    ExtraIcon    string             `json:"extraIcon,omitempty"`    // 菜单名称右侧的额外图标
    ShowLink     bool               `json:"showLink"`               // 是否显示该菜单
    ShowParent   bool               `json:"showParent"`             // 是否显示父级菜单
    Auths        []string           `json:"auths,omitempty"`        // 按钮级别权限设置
    KeepAlive    bool               `json:"keepAlive"`              // 是否缓存该路由页面（开启后，会保存该页面的整体状态，刷新后会清空状态）
    FrameSrc     string             `json:"frameSrc,omitempty"`     // 需要内嵌的iframe链接地址
    FrameLoading bool               `json:"frameLoading,omitempty"` // 内嵌的iframe页面是否开启首次加载动画
    Transition   PureMenuTransition `json:"transition,omitempty"`   // 页面动画
    HiddenTag    bool               `json:"hiddenTag"`              // 当前菜单名称或自定义信息禁止添加到标签页
    ActivePath   string             `json:"activePath,omitempty"`   // 将某个菜单激活
    FixedTag     bool               `json:"fixedTag"`               // 是否固定标签页
}

// PureMenuTransition pure admin 菜单页面动画
type PureMenuTransition struct {
    Name            string `json:"name,omitempty"`            // 当前页面动画
    EnterTransition string `json:"enterTransition,omitempty"` // 当前页面进场动画
    LeaveTransition string `json:"leaveTransition,omitempty"` // 当前页面离场动画
}
