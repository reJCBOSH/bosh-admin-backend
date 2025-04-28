package initialize

import (
	"bosh-admin/core/db"
	"bosh-admin/global"
)

// InitDB 初始化数据库
func InitDB() {
	global.GormDB = db.ConnectPgsql(global.Config.Pgsql)
	if global.GormDB == nil {
		panic("连接数据库失败")
	}
}
