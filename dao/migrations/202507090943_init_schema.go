package migrations

import (
    "time"

    "bosh-admin/core/log"
    "bosh-admin/dao"
    "bosh-admin/dao/model"
    "bosh-admin/global"
    "bosh-admin/utils"

    "github.com/go-gormigrate/gormigrate/v2"
    "gorm.io/gorm"
)

var InitSchema = &gormigrate.Migration{
    ID: "202507090943_init_schema",
    Migrate: func(tx *gorm.DB) error {
        err := tx.Migrator().AutoMigrate(
            &model.SysBlackJwt{},
            &model.SysDept{},
            &model.SysMenu{},
            &model.SysRole{},
            &model.SysRoleDept{},
            &model.SysRoleMenu{},
            &model.SysUser{},
            &model.SysLoginRecord{},
            &model.SysOperationRecord{},
        )
        if err != nil {
            log.Error("初始化数据表失败", err)
            return err
        }
        dept := model.SysDept{
            DeptName:     "系统管理",
            DeptCode:     global.SystemAdmin,
            Remark:       "系统管理",
            Status:       1,
            DisplayOrder: 9999,
        }
        err = tx.Create(&dept).Error
        if err != nil {
            log.Error("初始化部门数据失败", err)
            return err
        }
        role := model.SysRole{
            RoleName: "超级管理员",
            RoleCode: global.SuperAdmin,
            Status:   1,
            Remark:   "超级管理员",
            DataAuth: 1,
        }
        err = tx.Create(&role).Error
        if err != nil {
            log.Error("初始化角色数据表失败", err)
            return err
        }
        defaultPwd, _ := utils.BcryptHash(global.DefaultPassword)
        user := model.SysUser{
            Username:     global.SuperAdmin,
            Password:     defaultPwd,
            PwdUpdatedAt: dao.CustomTime(time.Now()),
            Nickname:     "超级管理员",
            Gender:       0,
            Introduce:    "行天之道，总司一切",
            Status:       1,
            RoleId:       role.Id,
            DeptId:       dept.Id,
        }
        err = tx.Create(&user).Error
        if err != nil {
            log.Error("初始化用户数据表失败", err)
            return err
        }
        return nil
    },
    Rollback: func(tx *gorm.DB) error {
        return tx.Migrator().DropTable(
            &model.SysBlackJwt{},
            &model.SysDept{},
            &model.SysMenu{},
            &model.SysRole{},
            &model.SysRoleDept{},
            &model.SysRoleMenu{},
            &model.SysUser{},
        )
    },
}
