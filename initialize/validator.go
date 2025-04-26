package initialize

import (
	"fmt"
	"reflect"
	"strings"

	"bosh-admin/global"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// InitValidator 初始化校验器
func InitValidator() {
	validate := validator.New()
	// 注册获取 json tag 的函数
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	// 注册中文翻译器
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	global.Trans, _ = uni.GetTranslator("zh")
	err := zhTranslations.RegisterDefaultTranslations(validate, global.Trans)
	if err != nil {
		panic(fmt.Errorf("注册中文翻译器失败: %s \n", err.Error()))
	}

	// 注册自定义检验方法

	// 注册自定义错误提示
}

// registerTranslator 为自定义校验字段添加翻译功能
func registerTranslator(tag, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
