package service

import (
    "bosh-admin/dao"
    "bosh-admin/dao/model"
)

type SysOperationRecordSvc struct{}

func NewSysOperationRecordSvc() *SysOperationRecordSvc {
    return &SysOperationRecordSvc{}
}

func (svc *SysOperationRecordSvc) GetOperationRecordList(username, method, path string, status int, requestIp string, startTime, endTime string, pageNo, pageSize int) ([]model.SysOperationRecord, int64, error) {
    s := dao.NewStatement()
    s.Select("id,created_at,uid,username,method,path,status,latency,request_ip,request_region,request_os,request_browser")
    if username != "" {
        s.Where("username like ?", "%"+username+"%")
    }
    if method != "" {
        s.Where("method = ?", method)
    }
    if path != "" {
        s.Where("path like ?", "%"+path+"%")
    }
    if status != 0 {
        s.Where("status = ?", status)
    }
    if requestIp != "" {
        s.Where("request_ip like ?", "%"+requestIp+"%")
    }
    if startTime != "" && endTime != "" {
        s.Where("created_at BETWEEN ? AND ?", startTime, endTime)
    }
    s.Pagination(pageNo, pageSize)
    return dao.QueryList[model.SysOperationRecord](s)
}

func (svc *SysOperationRecordSvc) GetOperationRecordById(id uint) (model.SysOperationRecord, error) {
    return dao.QueryById[model.SysOperationRecord](id)
}

func (svc *SysOperationRecordSvc) DelOperationRecord(id uint) error {
    return dao.DelById[model.SysOperationRecord](id)
}

func (svc *SysOperationRecordSvc) BatchDelOperationRecord(ids []uint) error {
    return dao.DelByIds[model.SysOperationRecord](ids)
}
