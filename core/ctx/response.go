package ctx

import (
	"net/http"

	"bosh-admin/dao/dto"
	"bosh-admin/exception"
	"bosh-admin/global"
)

// Response 响应
func (c *Context) Response(success bool, data any, msg string) {
	c.JSON(http.StatusOK, dto.Response{
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
	c.Response(dto.SUCCESS, nil, exceptionMsg)
}

// SuccessWithData 成功数据响应
func (c *Context) SuccessWithData(data any) {
	c.Response(dto.SUCCESS, data, exception.Success)
}

// SuccessWithList 成功列表响应
func (c *Context) SuccessWithList(list any, total int64) {
	c.Response(dto.SUCCESS, dto.ListData{
		List:  list,
		Total: total,
	}, exception.Success)
}

// SuccessWithDetail 成功详情响应
func (c *Context) SuccessWithDetail(data any, msg string) {
	c.Response(dto.SUCCESS, data, msg)
}

// Fail 失败响应
func (c *Context) Fail(msg ...string) {
	var exceptionMsg = exception.ServerError
	if len(msg) > 0 {
		c.Response(dto.FAIL, nil, msg[0])
	}
	c.Response(dto.FAIL, nil, exceptionMsg)
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
	c.JSON(http.StatusUnauthorized, dto.Response{
		Success: dto.FAIL,
		Data:    nil,
		Msg:     msg,
	})
}
