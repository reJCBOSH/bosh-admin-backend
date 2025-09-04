package dto

import "bosh-admin/dao"

type GetApiListReq struct {
	Pagination
	ApiName    string `json:"apiName" form:"apiName"`
	ApiGroup   string `json:"apiGroup" form:"apiGroup"`
	ApiPath    string `json:"apiPath" form:"apiPath"`
	ApiMethod  string `json:"apiMethod" form:"apiMethod"`
	IsRequired *int   `json:"isRequired" form:"isRequired" validate:"omitempty,oneof=0 1"`
}

type AddApiReq struct {
	dao.AddBasicModel
	ApiName    string `json:"apiName" validate:"required"`
	ApiGroup   string `json:"apiGroup" validate:"required"`
	ApiPath    string `json:"apiPath" validate:"required"`
	ApiMethod  string `json:"apiMethod" validate:"required,oneof=GET POST"`
	ApiDesc    string `json:"apiDesc"`
	IsRequired int    `json:"isRequired" validate:"oneof=0 1"`
}

type EditApiReq struct {
	dao.EditBasicModel
	ApiName    string `json:"apiName" validate:"required"`
	ApiGroup   string `json:"apiGroup" validate:"required"`
	ApiPath    string `json:"apiPath" validate:"required"`
	ApiMethod  string `json:"apiMethod" validate:"required,oneof=GET POST"`
	ApiDesc    string `json:"apiDesc"`
	IsRequired int    `json:"isRequired" validate:"oneof=0 1"`
}
