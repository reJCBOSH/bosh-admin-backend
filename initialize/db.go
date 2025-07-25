package initialize

import (
    "bosh-admin/core/db"
    "bosh-admin/core/log"
    "bosh-admin/dao/migrations"
    "bosh-admin/global"
)

// InitDB 初始化数据库
func InitDB() {
    switch global.Config.Server.Database {
    case "mysql":
        global.GormDB = db.ConnectMysql(global.Config.Mysql)
    case "pgsql":
        global.GormDB = db.ConnectPgsql(global.Config.Pgsql)
    }
    if global.GormDB == nil {
        panic("连接数据库失败")
    }
    log.Info("数据库连接成功")
    err := migrations.MigrateDatabase()
    if err != nil {
        panic("数据库迁移失败")
    }
    log.Info("数据库迁移成功")
}
