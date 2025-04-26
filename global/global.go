package global

import (
	"bosh-admin/config"

	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
)

// 环境常量
const (
	DEV  = "dev"
	UAT  = "uat"
	PROD = "prod"
)

var (
	Config config.Config      // 配置
	Logger *zap.SugaredLogger // 日志
	Trans  ut.Translator      // 翻译器
)
