package dto

type GetLoginRecordListRequest struct {
    Pagination
    Username  string `json:"username" form:"username"`
    StartDate string `json:"startDate" form:"startDate"`
    EndDate   string `json:"endDate" form:"endDate"`
}
