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
            &model.Resource{},
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
        menus := []model.SysMenu{
            {Path: "/system", Name: "System", ParentId: 0, Title: "系统管理", Icon: "ri:settings-3-line", ShowLink: true},
            {Path: "/system/user", Name: "SystemUser", Component: "system/user/index", ParentId: 1, Title: "用户管理", Icon: "ri:admin-line", ShowLink: true, KeepAlive: true},
            {Path: "/system/role", Name: "SystemRole", Component: "system/role/index", ParentId: 1, Title: "角色管理", Icon: "ri:admin-fill", ShowLink: true, KeepAlive: true},
            {Path: "/system/menu", Name: "SystemMenu", Component: "system/menu/index", ParentId: 1, Title: "菜单管理", Icon: "ep:menu", ShowLink: true, KeepAlive: true},
            {Path: "/system/dept", Name: "SystemDept", Component: "system/dept/index", ParentId: 1, Title: "部门管理", Icon: "ri:git-branch-line", ShowLink: true, KeepAlive: true},
            {Path: "/monitor", Name: "Monitor", ParentId: 0, Title: "系统监控", Icon: "ep:monitor", ShowLink: true},
            {Path: "/monitor/loginRecord", Name: "LoginRecord", Component: "monitor/loginRecord/index", ParentId: 6, Title: "登录日志", Icon: "ri:window-line", ShowLink: true, KeepAlive: true},
            {Path: "/monitor/operationRecord", Name: "OperationRecord", Component: "monitor/operationRecord/index", ParentId: 6, Title: "操作日志", Icon: "ri:history-fill", ShowLink: true, KeepAlive: true},
            // 用户按钮
            {ParentId: 2, MenuType: 3, Title: "新增用户", AuthCode: "sysUser:add"},
            {ParentId: 2, MenuType: 3, Title: "修改用户", AuthCode: "sysUser:edit"},
            {ParentId: 2, MenuType: 3, Title: "删除用户", AuthCode: "sysUser:del"},
            {ParentId: 2, MenuType: 3, Title: "设置用户状态", AuthCode: "sysUser:status"},
            {ParentId: 2, MenuType: 3, Title: "重置密码", AuthCode: "sysUser:resetPassword"},
            // 角色按钮
            {ParentId: 3, MenuType: 3, Title: "新增角色", AuthCode: "sysRole:add"},
            {ParentId: 3, MenuType: 3, Title: "修改角色", AuthCode: "sysRole:edit"},
            {ParentId: 3, MenuType: 3, Title: "删除角色", AuthCode: "sysRole:del"},
            {ParentId: 3, MenuType: 3, Title: "设置角色状态", AuthCode: "sysRole:status"},
            {ParentId: 3, MenuType: 3, Title: "设置菜单权限", AuthCode: "sysRole:menuAuth"},
            {ParentId: 3, MenuType: 3, Title: "设置数据权限", AuthCode: "sysRole:dataAuth"},
            // 菜单按钮
            {ParentId: 4, MenuType: 3, Title: "新增菜单", AuthCode: "sysMenu:add"},
            {ParentId: 4, MenuType: 3, Title: "修改菜单", AuthCode: "sysMenu:edit"},
            {ParentId: 4, MenuType: 3, Title: "删除菜单", AuthCode: "sysMenu:del"},
            // 部门按钮
            {ParentId: 5, MenuType: 3, Title: "新增部门", AuthCode: "sysDept:add"},
            {ParentId: 5, MenuType: 3, Title: "修改部门", AuthCode: "sysDept:edit"},
            {ParentId: 5, MenuType: 3, Title: "删除部门", AuthCode: "sysDept:del"},
            // 登录日志按钮
            {ParentId: 7, MenuType: 3, Title: "删除日志", AuthCode: "sysLoginRecord:del"},
            {ParentId: 7, MenuType: 3, Title: "批量删除日志", AuthCode: "sysLoginRecord:batchDel"},
            // 操作日志按钮
            {ParentId: 8, MenuType: 3, Title: "查看详情", AuthCode: "sysOperationRecord:view"},
            {ParentId: 8, MenuType: 3, Title: "删除日志", AuthCode: "sysOperationRecord:del"},
            {ParentId: 8, MenuType: 3, Title: "批量删除日志", AuthCode: "sysOperationRecord:batchDel"},
        }
        err = tx.Create(&menus).Error
        if err != nil {
            log.Error("初始化菜单数据表失败", err)
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
            &model.SysLoginRecord{},
            &model.SysOperationRecord{},
            &model.Resource{},
        )
    },
}
