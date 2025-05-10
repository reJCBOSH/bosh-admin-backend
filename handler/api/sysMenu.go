package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/dao/dto"
	"bosh-admin/dao/model"
	"bosh-admin/service"
)

type SysMenuHandler struct {
	svc service.SysMenuSvc
}

func NewSysMenuHandler() *SysMenuHandler {
	return &SysMenuHandler{}
}

func (h *SysMenuHandler) GetMenuTree(c *ctx.Context) {
	menuTree, err := h.svc.GetMenuTree()
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(menuTree)
}

func (h *SysMenuHandler) GetMenuList(c *ctx.Context) {
	var req dto.GetMenuListRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	menus, total, err := h.svc.GetMenus(req.Title, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	var list []dto.MenuItem
	for _, menu := range menus {
		list = append(list, dto.MenuItem{
			BasicModel: menu.BasicModel,
			MenuBasicItem: dto.MenuBasicItem{
				Path:            menu.Path,
				Name:            menu.Name,
				Redirect:        menu.Redirect,
				Component:       menu.Component,
				ParentId:        menu.ParentId,
				MenuType:        menu.MenuType,
				Title:           menu.Title,
				Icon:            menu.Icon,
				DisplayOrder:    menu.DisplayOrder,
				ExtraIcon:       menu.ExtraIcon,
				Transition:      menu.Transition,
				EnterTransition: menu.EnterTransition,
				LeaveTransition: menu.LeaveTransition,
				ActivePath:      menu.ActivePath,
				AuthCode:        menu.AuthCode,
				FrameSrc:        menu.FrameSrc,
				FrameLoading:    menu.FrameLoading,
				ShowLink:        menu.ShowLink,
				ShowParent:      menu.ShowParent,
				KeepAlive:       menu.KeepAlive,
				HiddenTag:       menu.HiddenTag,
				FixedTag:        menu.FixedTag,
			},
		})
	}
	c.SuccessWithList(list, total)
}

func (h *SysMenuHandler) GetMenuInfo(c *ctx.Context) {
	var req dto.IdRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	menu, err := h.svc.GetMenuById(req.Id)
	if c.HandlerError(err) {
		return
	}
	data := dto.MenuItem{
		BasicModel: menu.BasicModel,
		MenuBasicItem: dto.MenuBasicItem{
			Path:            menu.Path,
			Name:            menu.Name,
			Redirect:        menu.Redirect,
			Component:       menu.Component,
			ParentId:        menu.ParentId,
			MenuType:        menu.MenuType,
			Title:           menu.Title,
			Icon:            menu.Icon,
			DisplayOrder:    menu.DisplayOrder,
			ExtraIcon:       menu.ExtraIcon,
			Transition:      menu.Transition,
			EnterTransition: menu.EnterTransition,
			LeaveTransition: menu.LeaveTransition,
			ActivePath:      menu.ActivePath,
			AuthCode:        menu.AuthCode,
			FrameSrc:        menu.FrameSrc,
			FrameLoading:    menu.FrameLoading,
			ShowLink:        menu.ShowLink,
			ShowParent:      menu.ShowParent,
			KeepAlive:       menu.KeepAlive,
			HiddenTag:       menu.HiddenTag,
			FixedTag:        menu.FixedTag,
		},
	}
	c.SuccessWithData(data)
}

func (h *SysMenuHandler) AddMenu(c *ctx.Context) {
	var req dto.AddMenuRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	menu := model.SysMenu{
		Path:            req.Path,
		Name:            req.Name,
		Redirect:        req.Redirect,
		Component:       req.Component,
		ParentId:        req.ParentId,
		MenuType:        req.MenuType,
		Title:           req.Title,
		Icon:            req.Icon,
		DisplayOrder:    req.DisplayOrder,
		ExtraIcon:       req.ExtraIcon,
		Transition:      req.Transition,
		EnterTransition: req.EnterTransition,
		LeaveTransition: req.LeaveTransition,
		ActivePath:      req.ActivePath,
		AuthCode:        req.AuthCode,
		FrameSrc:        req.FrameSrc,
		FrameLoading:    req.FrameLoading,
		ShowLink:        req.ShowLink,
		ShowParent:      req.ShowParent,
		KeepAlive:       req.KeepAlive,
		HiddenTag:       req.HiddenTag,
		FixedTag:        req.FixedTag,
	}
	err = h.svc.AddMenu(menu)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysMenuHandler) EditMenu(c *ctx.Context) {
	var req dto.EditMenuRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}

}

func (h *SysMenuHandler) DelMenu(c *ctx.Context) {
	var req dto.IdRequest
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.DelMenu(req.Id)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysMenuHandler) GetAsyncRoutes(c *ctx.Context) {

}
