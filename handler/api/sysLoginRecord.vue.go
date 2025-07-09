package api

import (
    "bosh-admin/core/ctx"
    "bosh-admin/dao/dto"
    "bosh-admin/service"
)

type SysLoginRecord struct {
    svc *service.SysLoginRecordSvc
}

func NewSysLoginRecordHandler() *SysLoginRecord {
    return &SysLoginRecord{
        svc: service.NewSysLoginRecordSvc(),
    }
}

func (h *SysLoginRecord) GetLoginRecordList(c *ctx.Context) {
    var req dto.GetLoginRecordListRequest
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    records, total, err := h.svc.GetLoginRecordList(req.Username, req.StartDate, req.EndDate, req.PageNo, req.PageSize)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithList(records, total)
}

func (h *SysLoginRecord) DelLoginRecord(c *ctx.Context) {
    var req dto.IdRequest
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.DelLoginRecord(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysLoginRecord) BatchDelLoginRecord(c *ctx.Context) {
    var req dto.IdsRequest
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.BatchDelLoginRecord(req.Ids)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}
