package handler

import (
	"bosh-admin/core/ctx"
	"bosh-admin/dao/dto"
	"bosh-admin/dao/model"
	"bosh-admin/service"
)

type SysApiHandler struct {
	svc       *service.SysApiSvc
	casbinSvc *service.CasbinSvc
}

func NewSysApiHandler() *SysApiHandler {
	return &SysApiHandler{
		svc:       service.NewSysApiSvc(),
		casbinSvc: service.NewCasbinSvc(),
	}
}

func (h *SysApiHandler) GetApiList(c *ctx.Context) {
	var req dto.GetApiListReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	list, total, err := h.svc.GetApiList(req.ApiName, req.ApiGroup, req.ApiPath, req.ApiMethod, req.IsRequired, req.PageNo, req.PageSize)
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithList(list, total)
}

func (h *SysApiHandler) GetApiGroups(c *ctx.Context) {
	data, err := h.svc.GetApiGroups()
	if c.HandlerError(err) {
		return
	}
	c.SuccessWithData(data)
}

func (h *SysApiHandler) AddApi(c *ctx.Context) {
	var req dto.AddApiReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	err = h.svc.AddApi(req)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysApiHandler) EditApi(c *ctx.Context) {
	var req dto.EditApiReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	var api *model.SysApi
	api, err = h.svc.EditApi(req)
	if c.HandlerError(err) {
		return
	}
	err = h.casbinSvc.UpdateCasbinApi(api.ApiPath, req.ApiPath, api.ApiMethod, req.ApiMethod)
	if c.HandlerError(err) {
		return
	}
	c.Success()
}

func (h *SysApiHandler) DelApi(c *ctx.Context) {
	var req dto.IdReq
	msg, err := c.ValidateParams(&req)
	if c.HandlerError(err, msg) {
		return
	}
	var api *model.SysApi
	api, err = h.svc.DelApi(req.Id)
	if c.HandlerError(err) {
		return
	}
	err = h.casbinSvc.RemoveCasbin(api.ApiPath, api.ApiMethod)
	c.Success()
}
