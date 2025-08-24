package handler

import (
    "bosh-admin/core/ctx"
    "bosh-admin/dao/dto"
    "bosh-admin/service"
)

type SysOperationRecord struct {
    svc *service.SysOperationRecordSvc
}

func NewSysOperationRecordHandler() *SysOperationRecord {
    return &SysOperationRecord{
        svc: service.NewSysOperationRecordSvc(),
    }
}

func (h *SysOperationRecord) GetOperationRecordList(c *ctx.Context) {
    var req dto.GetOperationRecordListReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    records, total, err := h.svc.GetOperationRecordList(req.Username, req.Method, req.Path, req.Status, req.RequestIP, req.StartTime, req.EndTime, req.PageNo, req.PageSize)
    if c.HandlerError(err) {
        return
    }
    var list []dto.OperationRecordListItem
    for _, record := range records {
        list = append(list, dto.OperationRecordListItem{
            Id:             record.Id,
            CreatedAt:      record.CreatedAt.String(),
            Username:       record.Username,
            Method:         record.Method,
            Path:           record.Path,
            Status:         record.Status,
            Latency:        record.Latency,
            RequestIP:      record.RequestIP,
            RequestRegion:  record.RequestRegion,
            RequestOS:      record.RequestOS,
            RequestBrowser: record.RequestBrowser,
        })
    }
    c.SuccessWithList(list, total)
}

func (h *SysOperationRecord) GetOperationRecordInfo(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    data, err := h.svc.GetOperationRecordById(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.SuccessWithData(data)
}

func (h *SysOperationRecord) DelOperationRecord(c *ctx.Context) {
    var req dto.IdReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.DelOperationRecord(req.Id)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}

func (h *SysOperationRecord) BatchDelOperationRecord(c *ctx.Context) {
    var req dto.IdsReq
    msg, err := c.ValidateParams(&req)
    if c.HandlerError(err, msg) {
        return
    }
    err = h.svc.BatchDelOperationRecord(req.Ids)
    if c.HandlerError(err) {
        return
    }
    c.Success()
}
