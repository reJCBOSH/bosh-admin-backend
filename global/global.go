package global

import (
	"bosh-admin/config"

	"go.uber.org/zap"
)

var (
	Config config.Config      // 配置
	Logger *zap.SugaredLogger // 日志
)
