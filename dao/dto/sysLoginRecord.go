package dto

type GetLoginRecordListReq struct {
    Pagination
    Username  string `json:"username" form:"username"`
    StartTime string `json:"startTime" form:"startTime"`
    EndTime   string `json:"endTime" form:"endTime"`
}
