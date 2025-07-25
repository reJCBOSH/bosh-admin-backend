package dto

// IdReq id请求
type IdReq struct {
    Id uint `json:"id" form:"id" validate:"required,min=1"` // id
}

// IdsReq ids请求
type IdsReq struct {
    Ids []uint `json:"ids" form:"ids" validate:"required,gt=0,dive,min=1"` // ids
}

// Pagination 分页
type Pagination struct {
    PageNo   int `json:"pageNo" form:"pageNo" validate:"required,min=-1,ne=0"`                       // 页码
    PageSize int `json:"pageSize" form:"pageSize" validate:"required_unless=PageNo -1|gt=0,max=100"` // 每页数量
}

// OrderBy 排序
type OrderBy struct {
    Field string `json:"field" form:"field" validate:"omitempty"`              // 排序字段
    Rule  string `json:"rule" form:"rule" validate:"omitempty,oneof=ASC DESC"` // 排序规则
}
