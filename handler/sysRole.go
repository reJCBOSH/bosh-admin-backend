package handler

import (
    "bosh-admin/core/ctx"
    "bosh-admin/dao/dto"
    "bosh-admin/service"
)

type SysRoleHandler struct {
    svc    *service.SysRoleSvc
    jwtSvc *service.JWTSvc
}

func NewSysRoleHandler() *SysRoleHandler {
    return &SysRoleHandler{
        svc:    service.NewSysRoleSvc(),
        jwtSvc: service.NewJWTSvc(),
    }
}

func (h *SysRoleHandler) GetRoleList(c *ctx.Context) {
    var req dto.GetRoleListReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    list, total, err := h.svc.GetRoleList(req.RoleName, req.RoleCode, req.Status, req.PageNo, req.PageSize)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithList(list, total)
}

func (h *SysRoleHandler) GetRoleInfo(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    role, err := h.svc.GetRoleById(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(role)
}

func (h *SysRoleHandler) AddRole(c *ctx.Context) {
    var req dto.AddRoleReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.AddRole(req)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysRoleHandler) EditRole(c *ctx.Context) {
    var req dto.EditRoleReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.EditRole(req)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysRoleHandler) DelRole(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.DelRole(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysRoleHandler) GetRoleMenu(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    menus, err := h.svc.GetRoleMenu(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(menus)
}

func (h *SysRoleHandler) GetRoleMenuIds(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    data, err := h.svc.GetRoleMenuIds(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(data)
}

func (h *SysRoleHandler) SetRoleMenuAuth(c *ctx.Context) {
    var req dto.SetRoleMenuAuthReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.SetRoleMenuAuth(req.RoleId, req.MenuIds)
    if c.HandlerError(err) {
        return
    }
    // 判断是否同一角色
    userClaims := h.jwtSvc.GetUserClaims(c)
    c.SuccessWithData(req.RoleId == userClaims.RoleId)
}

func (h *SysRoleHandler) GetRoleDeptIds(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    data, err := h.svc.GetRoleDeptIds(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(data)
}

func (h *SysRoleHandler) SetRoleDataAuth(c *ctx.Context) {
    var req dto.SetRoleDataAuthReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.SetRoleDataAuth(req.RoleId, req.DataAuth, req.DeptIds)
    if c.HandlerError(err) {
        return
    }
    // 判断是否同一角色
    userClaims := h.jwtSvc.GetUserClaims(c)
    c.SuccessWithData(req.RoleId == userClaims.RoleId)
}

func (h *SysRoleHandler) SetRoleStatus(c *ctx.Context) {
    var req dto.SetRoleStatusReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    // 判断是否同一角色
    userClaims := h.jwtSvc.GetUserClaims(c)
    err = h.svc.SetRoleStatus(userClaims.RoleId, req.RoleId, req.Status)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}
