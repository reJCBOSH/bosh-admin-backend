package db

import (
	"fmt"

	"bosh-admin/config"
	"bosh-admin/core/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// ConnectPgsql 连接pgsql
func ConnectPgsql(p config.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", p.Host, p.User, p.Password, p.Dbname, p.Port, p.Config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		Logger:                                   log.CustomGormLogger("_pgsql"),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	return db
}
