package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/dao/dto"
    "bosh-admin/service"
)

type SysDeptHandler struct {
    svc *service.SysDeptSvc
}

func NewSysDeptHandler() *SysDeptHandler {
    return &SysDeptHandler{
        svc: service.NewSysDeptSvc(),
    }
}

func (h *SysDeptHandler) GetDeptTree(c *ctx.Context) {
    deptTree, err := h.svc.GetDeptTree()
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(deptTree)
}

func (h *SysDeptHandler) GetDeptList(c *ctx.Context) {
    var req dto.GetDeptListReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    list, total, err := h.svc.GetDeptList(req.DeptName, req.DeptCode, req.Status, req.PageNo, req.PageSize)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithList(list, total)
}

func (h *SysDeptHandler) GetDeptInfo(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    dept, err := h.svc.GetDeptById(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(dept)
}

func (h *SysDeptHandler) AddDept(c *ctx.Context) {
    var req dto.AddDeptReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.AddDept(req)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysDeptHandler) EditDept(c *ctx.Context) {
    var req dto.EditDeptReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.EditDept(req)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysDeptHandler) DelDept(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.DelDept(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}
