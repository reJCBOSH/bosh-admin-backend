package service

import (
    "bosh-admin/core/exception"
    "bosh-admin/dao"
    "bosh-admin/dao/dto"
    "bosh-admin/dao/model"
    "bosh-admin/global"
)

type SysRoleSvc struct {
}

func NewSysRoleSvc() *SysRoleSvc {
    return &SysRoleSvc{}
}

func (svc *SysRoleSvc) GetRoleList(roleName, roleCode string, status *int, pageNo, pageSize int) ([]model.SysRole, int64, error) {
    s := dao.NewStatement()
    if roleName != "" {
        s.Where("role_name LIKE ?", "%"+roleName+"%")
    }
    if roleCode != "" {
        s.Where("role_code LIKE ?", "%"+roleCode+"%")
    }
    if status != nil {
        s.Where("status = ?", *status)
    }
    s.Pagination(pageNo, pageSize)
    return dao.QueryList[model.SysRole](s)
}

func (svc *SysRoleSvc) GetRoleById(id any) (model.SysRole, error) {
    return dao.QueryById[model.SysRole](id)
}

func (svc *SysRoleSvc) AddRole(role dto.AddRoleReq) error {
    s := dao.NewStatement()
    s.Where("role_code = ?", role.RoleCode)
    duplicateCode, err := dao.Count[model.SysRole](s)
    if err != nil {
        return err
    }
    if duplicateCode > 0 {
        return exception.NewException("角色标识已存在")
    }
    return dao.Create(&role, "sys_role")
}

func (svc *SysRoleSvc) EditRole(role dto.EditRoleReq) error {
    return dao.Updates(&role, "sys_role")
}

func (svc *SysRoleSvc) DelRole(id any) error {
    role, err := dao.QueryById[model.SysRole](id)
    if err != nil {
        return err
    }
    if role.RoleCode == global.SuperAdmin {
        return exception.NewException("无法删除超级管理员角色")
    }
    s := dao.NewStatement()
    s.Where("role_id = ?", id)
    bindUser, err := dao.Count[model.SysUser](s)
    if err != nil {
        return err
    }
    if bindUser > 0 {
        return exception.NewException("存在绑定用户，无法删除")
    }
    tx := dao.Begin()
    // 删除角色
    if err = tx.Delete(&model.SysRole{}, id).Error; err != nil {
        tx.Rollback()
        return err
    }
    // 删除角色-菜单关联
    if err = tx.Where("role_id = ?", id).Delete(&model.SysRoleMenu{}).Error; err != nil {
        tx.Rollback()
        return err
    }
    // 删除角色-部门关联
    if err = tx.Where("role_id = ?", id).Delete(&model.SysRoleDept{}).Error; err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return nil
}

func (svc *SysRoleSvc) GetRoleMenu(roleId any) ([]model.SysMenu, error) {
    //role, err := dao.QueryById[model.SysRole](roleId)
    //if err != nil {
    //    return nil, err
    //}
    s := dao.NewStatement()
    //if role.RoleCode != global.SuperAdmin {
    //    s.Where("create_by != ?", 0)
    //}
    s.OrderBy("display_order DESC,id ASC")
    data, _, err := dao.QueryList[model.SysMenu](s)
    return data, err
}

func (svc *SysRoleSvc) GetRoleMenuIds(roleId any) ([]uint, error) {
    role, err := dao.QueryById[model.SysRole](roleId)
    if err != nil {
        return nil, err
    }
    var ids []uint
    err = dao.GormDB().Model(&model.SysRoleMenu{}).Where("role_id = ?", role.Id).Pluck("menu_id", &ids).Error
    return ids, err
}

func (svc *SysRoleSvc) SetRoleMenuAuth(roleId uint, menuIds []uint) error {
    role, err := dao.QueryById[model.SysRole](roleId)
    if err != nil {
        return err
    }
    if role.RoleCode == global.SuperAdmin {
        return exception.NewException("无法设置超级管理员角色菜单权限")
    }
    s := dao.NewStatement()
    s.Where("id IN ?", menuIds)
    menuNum, err := dao.Count[model.SysMenu](s)
    if err != nil {
        return err
    }
    if menuNum != int64(len(menuIds)) {
        return exception.NewException("菜单权限数据错误")
    }
    var roleMenus []model.SysRoleMenu
    for _, v := range menuIds {
        roleMenus = append(roleMenus, model.SysRoleMenu{
            RoleId: roleId,
            MenuId: v,
        })
    }
    tx := dao.Begin()
    if err = tx.Where("role_id = ?", roleId).Delete(&model.SysRoleMenu{}).Error; err != nil {
        tx.Rollback()
        return err
    }
    if err = tx.Save(&roleMenus).Error; err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return nil
}

func (svc *SysRoleSvc) GetRoleDeptIds(roleId any) ([]uint, error) {
    role, err := dao.QueryById[model.SysRole](roleId)
    if err != nil {
        return nil, err
    }
    var ids []uint
    err = dao.GormDB().Model(&model.SysRoleDept{}).Where("role_id = ?", role.Id).Pluck("dept_id", &ids).Error
    return ids, err
}

func (svc *SysRoleSvc) SetRoleDataAuth(roleId uint, dataAuth int, deptIds []uint) error {
    if dataAuth == 5 && len(deptIds) == 0 {
        return exception.NewException("请至少选择一个部门")
    }
    role, err := dao.QueryById[model.SysRole](roleId)
    if err != nil {
        return err
    }
    var roleDepts []model.SysRoleDept
    if dataAuth == 5 {
        s := dao.NewStatement()
        s.Where("id IN ?", deptIds)
        deptNum, err := dao.Count[model.SysDept](s)
        if err != nil {
            return err
        }
        if deptNum != int64(len(deptIds)) {
            return exception.NewException("部门数据错误")
        }
        for _, v := range deptIds {
            roleDepts = append(roleDepts, model.SysRoleDept{
                RoleId: role.Id,
                DeptId: v,
            })
        }
    }
    tx := dao.Begin()
    if role.DataAuth != dataAuth {
        if err = tx.Model(&model.SysRole{}).Where("id = ?", role.Id).Update("data_auth", dataAuth).Error; err != nil {
            tx.Rollback()
            return err
        }
    }
    if dataAuth == 5 {
        if err = tx.Where("role_id = ?", role.Id).Delete(&model.SysRoleDept{}).Error; err != nil {
            tx.Rollback()
            return err
        }
        if err = tx.Save(&roleDepts).Error; err != nil {
            tx.Rollback()
            return err
        }
    }
    tx.Commit()
    return nil
}

func (svc *SysRoleSvc) SetRoleStatus(currentRoleId, roleId uint, status int) error {
    role, err := dao.QueryById[model.SysRole](roleId)
    if err != nil {
        return err
    }
    if role.RoleCode == global.SuperAdmin {
        return exception.NewException("无法修改超级管理员角色状态")
    }
    if role.Status == status {
        return exception.NewException("角色状态未改变")
    }
    if status == 1 {
        // 判断是否分配菜单权限
        s := dao.NewStatement()
        s.Where("role_id = ?", role.Id)
        menuNum, err := dao.Count[model.SysRoleMenu](s)
        if err != nil {
            return err
        }
        if menuNum == 0 {
            return exception.NewException("请先分配菜单权限")
        }
        // 判断是否分配数据权限
        if role.DataAuth == 0 {
            return exception.NewException("请先分配数据权限")
        }
        // TODO 判断是否分配api权限
    } else {
        if currentRoleId == roleId {
            return exception.NewException("无法禁用当前操作员角色")
        }
    }
    role.Status = status
    return dao.Updates(&role)
}
