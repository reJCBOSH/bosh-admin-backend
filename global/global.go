package global

import (
    "bosh-admin/config"

    "github.com/gin-gonic/gin"
    ut "github.com/go-playground/universal-translator"
    "github.com/lionsoul2014/ip2region/binding/golang/xdb"
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
    Config      config.Config      // 配置
    Logger      *zap.SugaredLogger // 日志
    Trans       ut.Translator      // 翻译器
    GormDB      *gorm.DB           // gorm数据库
    Router      *gin.Engine        // 路由
    XdbSearcher *xdb.Searcher      // 全局xdb搜索器
)

const (
    SuperAdmin      = "SuperAdmin"
    SystemAdmin     = "SystemAdmin"
    DefaultPassword = "Ab112112."
)
