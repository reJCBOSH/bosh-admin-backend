package initialize

import (
    "bosh-admin/core/db"
    "bosh-admin/global"
)

// InitDB 初始化数据库
func InitDB() {
    global.GormDB = db.ConnectMysql(global.Config.Mysql)
    if global.GormDB == nil {
        panic("连接数据库失败")
    }
    err := db.MigrateDatabase()
    if err != nil {
        panic("数据库迁移失败")
    }
}
