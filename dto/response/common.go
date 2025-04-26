package response

const (
	SUCCESS = true
	FAIL    = false
)

// Response 响应结构体
type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
}

// ListData 列表数据
type ListData struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}
