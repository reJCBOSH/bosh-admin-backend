package ctx

import (
	"bosh-admin/dto/response"
	"bosh-admin/exception"
	"bosh-admin/global"
	"net/http"
)

// Response 响应
func (c *Context) Response(success bool, data any, msg string) {
	c.JSON(http.StatusOK, response.Response{
		Success: success,
		Data:    data,
		Msg:     msg,
	})
}

// Success 成功响应
func (c *Context) Success(msg ...string) {
	var exceptionMsg = exception.Success
	if len(msg) > 0 {
		exceptionMsg = msg[0]
	}
	c.Response(response.SUCCESS, nil, exceptionMsg)
}

// SuccessWithData 成功数据响应
func (c *Context) SuccessWithData(data any) {
	c.Response(response.SUCCESS, data, exception.Success)
}

// SuccessWithList 成功列表响应
func (c *Context) SuccessWithList(list any, total int64) {
	c.Response(response.SUCCESS, response.ListData{
		List:  list,
		Total: total,
	}, exception.Success)
}

// SuccessWithDetail 成功详情响应
func (c *Context) SuccessWithDetail(data any, msg string) {
	c.Response(response.SUCCESS, data, msg)
}

// Fail 失败响应
func (c *Context) Fail(msg ...string) {
	var exceptionMsg = exception.ServerError
	if len(msg) > 0 {
		c.Response(response.FAIL, nil, msg[0])
	}
	c.Response(response.FAIL, nil, exceptionMsg)
}

// HandlerError 错误处理
func (c *Context) HandlerError(err any, msg ...string) bool {
	if err != nil {
		global.Logger.Error(err)
		c.Fail(msg...)
		return true
	}
	return false
}

// UnAuthorized 未授权响应
func (c *Context) UnAuthorized(msg string) {
	c.JSON(http.StatusUnauthorized, response.Response{
		Success: response.FAIL,
		Data:    nil,
		Msg:     msg,
	})
}
