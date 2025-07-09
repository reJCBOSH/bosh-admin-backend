package service

import (
    "time"

    "bosh-admin/dao"
    "bosh-admin/dao/model"
    "bosh-admin/utils"

    ua "github.com/mssola/user_agent"
)

type SysLoginRecordSvc struct{}

func NewSysLoginRecordSvc() *SysLoginRecordSvc {
    return &SysLoginRecordSvc{}
}

func (svc *SysLoginRecordSvc) AddLoginRecord(uid uint, username, loginIP, userAgent string, loginStatus int) error {
    var record = model.SysLoginRecord{
        Uid:         uid,
        Username:    username,
        LoginIP:     loginIP,
        UserAgent:   userAgent,
        LoginStatus: loginStatus,
        LoginTime:   dao.CustomTime(time.Now().Local()),
    }
    record.LoginRegion = utils.IP2Region(loginIP)
    UA := ua.New(userAgent)
    record.LoginOS = UA.OS()
    record.LoginBrowser, _ = UA.Browser()
    return dao.Create(&record)
}

func (svc *SysLoginRecordSvc) GetLoginRecordList(username, startDate, endDate string, pageNo, pageSize int) ([]model.SysLoginRecord, int64, error) {
    s := dao.NewStatement()
    if username != "" {
        s.Where("username LIKE ?", "%"+username+"%")
    }
    if startDate != "" && endDate != "" {
        s.Where("login_time BETWEEN  ? AND ?", startDate+" 00:00:00", endDate+" 23:59:59")
    }
    s.Pagination(pageNo, pageSize)
    return dao.QueryList[model.SysLoginRecord](s)
}

func (svc *SysLoginRecordSvc) DelLoginRecord(id uint) error {
    return dao.DelById[model.SysLoginRecord](id)
}

func (svc *SysLoginRecordSvc) BatchDelLoginRecord(ids []uint) error {
    return dao.DelByIds[model.SysLoginRecord](ids)
}
