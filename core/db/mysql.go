package db

import (
	"fmt"

	"bosh-admin/config"
	"bosh-admin/core/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectMysql 连接mysql
func ConnectMysql(m config.Mysql) *gorm.DB {
	if m.Database == "" {
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.IP, m.Port, m.Database, m.Config)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		Logger:                                   log.CustomGormLogger("_mysql"),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db
}
