package global

import (
	"bosh-admin/config"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	GormDB *gorm.DB           // gorm数据库
	Router *gin.Engine        // 路由
)
