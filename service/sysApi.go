package service

import (
	"bosh-admin/core/exception"
	"bosh-admin/dao"
	"bosh-admin/dao/dto"
	"bosh-admin/dao/model"
	"errors"
)

type SysApiSvc struct{}

func NewSysApiSvc() *SysApiSvc {
	return &SysApiSvc{}
}

func (svc *SysApiSvc) GetApiList(apiName, apiGroup, apiPath, apiMethod string, isRequired *int, pageNo, pageSize int) ([]model.SysApi, int64, error) {
	s := dao.NewStatement()
	if apiName != "" {
		s.Where("api_name LIKE ?", "%"+apiName+"%")
	}
	if apiGroup != "" {
		s.Where("api_group = ?", apiGroup)
	}
	if apiPath != "" {
		s.Where("api_path LIKE ?", "%"+apiPath+"%")
	}
	if apiMethod != "" {
		s.Where("api_method = ?", apiMethod)
	}
	if isRequired != nil {
		s.Where("is_required = ?", *isRequired)
	}
	s.Pagination(pageNo, pageSize)
	return dao.QueryList[model.SysApi](s)
}

func (svc *SysApiSvc) GetApiById(id uint) (model.SysApi, error) {
	return dao.QueryById[model.SysApi](id)
}

func (svc *SysApiSvc) GetApiGroups() ([]string, error) {
	var groups []string
	err := dao.GormDB().Model(&model.SysApi{}).Distinct("api_group").Order("api_group ASC").Pluck("api_group", &groups).Error
	return groups, err
}

func (svc *SysApiSvc) AddApi(req dto.AddApiReq) error {
	s := dao.NewStatement()
	s.Where("api_path = ?", req.ApiPath)
	s.Where("api_method = ?", req.ApiMethod)
	count, err := dao.Count[model.SysApi](s)
	if err != nil {
		return err
	}
	if count > 0 {
		return exception.NewException("已存在相同路径和请求方法的接口")
	}
	return dao.Create(&req, "sys_api")
}

func (svc *SysApiSvc) EditApi(req dto.EditApiReq) (*model.SysApi, error) {
	api, err := dao.QueryById[model.SysApi](req.Id)
	if err != nil {
		if errors.Is(err, dao.NotFound) {
			return nil, exception.NewException("接口不存在")
		}
		return nil, err
	}
	err = dao.Updates(&req, "sys_api")
	if err != nil {
		return nil, err
	}
	return &api, nil
}

func (svc *SysApiSvc) DelApi(id uint) (*model.SysApi, error) {
	api, err := dao.QueryById[model.SysApi](id)
	if err != nil {
		if errors.Is(err, dao.NotFound) {
			return nil, exception.NewException("接口不存在")
		}
		return nil, err
	}
	err = dao.DelById[model.SysApi](id)
	if err != nil {
		return nil, err
	}
	return &api, nil
}
