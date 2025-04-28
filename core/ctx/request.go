package ctx

import (
	"errors"
	"fmt"
	"strings"

	"bosh-admin/exception"
	"bosh-admin/global"

	"github.com/go-playground/validator/v10"
)

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// ValidateParams 校验请求参数
func (c *Context) ValidateParams(req any) (string, any) {
	err := c.ShouldBind(req)
	if err != nil {
		// 获取validator.ValidationErrors类型的errors
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			return exception.ServerError, err
		}
		errsMap := removeTopStruct(errs.Translate(global.Trans))
		var errsArr []string
		for _, v := range errsMap {
			errsArr = append(errsArr, v)
		}
		return fmt.Sprintf("%s:%s", exception.ParamsError, strings.Join(errsArr, ",")), errsMap
	}
	return "", nil
}
