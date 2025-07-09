package migrations

import (
    "bosh-admin/core/log"
    "bosh-admin/dao/model"
    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
)

var CreateSysLoginRecordTable = &gormigrate.Migration{
    ID: "202507090950_create_sys_login_record_table",
    Migrate: func(tx *gorm.DB) error {
        err := tx.Migrator().CreateTable(&model.SysLoginRecord{})
        if err != nil {
            log.Error("创建表sys_login_record失败:", err)
        }
        return err
    },
    Rollback: func(tx *gorm.DB) error {
        err := tx.Migrator().DropTable(&model.SysLoginRecord{})
        if err != nil {
            log.Error("删除表sys_login_record失败:", err)
        }
        return err
    },
}
