package db

import (
	"fmt"

	"bosh-admin/config"
	"bosh-admin/core/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectPgsql 连接pgsql
func ConnectPgsql(p config.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", p.Host, p.User, p.Password, p.Dbname, p.Port, p.Config)
	pgsqlConfig := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 禁用prefer protocol
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
		Logger:                                   log.CustomGormLogger("_pgsql"),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	return db
}
