package api

import (
	"bosh-admin/core/ctx"
	"bosh-admin/dao/dto"
	"bosh-admin/service"
)

type SysMenuHandler struct {
	svc *service.SysMenuSvc
}

func NewSysMenuHandler() *SysMenuHandler {
	return &SysMenuHandler{
		svc: service.NewSysMenuSvc(),
	}
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
	menus, total, err := h.svc.GetMenuList(req.Title, req.PageNo, req.PageSize)
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
	c.SuccessWithData(dto.GetMenuListResponse{
		List:  list,
		Total: total,
	})
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
	err = h.svc.AddMenu(req)
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
	err = h.svc.EditMenu(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
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
