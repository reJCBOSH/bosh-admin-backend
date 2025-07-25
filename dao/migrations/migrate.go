package migrations

import (
    "bosh-admin/global"

    "github.com/go-gormigrate/gormigrate/v2"
)

func MigrateDatabase() error {
    migrationArr := []*gormigrate.Migration{
        InitSchema,
    }

    m := gormigrate.New(global.GormDB, gormigrate.DefaultOptions, migrationArr)
    return m.Migrate()
}
