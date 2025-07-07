package service

import (
    "bosh-admin/core/exception"
    "bosh-admin/dao"
    "bosh-admin/dao/dto"
    "bosh-admin/dao/model"
    "bosh-admin/global"
)

type SysDeptSvc struct{}

func NewSysDeptSvc() *SysDeptSvc {
    return &SysDeptSvc{}
}

func (svc *SysDeptSvc) GetDeptTree() (deptTree []model.SysDept, err error) {
    treeMap, err := queryDeptTreeMap()
    deptTree = treeMap[0]
    for i := 0; i < len(deptTree); i++ {
        err = getDeptChildrenList(&deptTree[i], treeMap)
    }
    return deptTree, err
}

// queryDeptTreeMap 查询部门树map
func queryDeptTreeMap() (map[uint][]model.SysDept, error) {
    var allDept []model.SysDept
    treeMap := make(map[uint][]model.SysDept)
    err := dao.GormDB().Order("display_order DESC").Find(&allDept).Error
    for _, v := range allDept {
        treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
    }
    return treeMap, err
}

// getDeptChildrenList 获取子部门列表
func getDeptChildrenList(dept *model.SysDept, treeMap map[uint][]model.SysDept) (err error) {
    dept.Children = treeMap[dept.Id]
    for i := 0; i < len(dept.Children); i++ {
        err = getDeptChildrenList(&dept.Children[i], treeMap)
    }
    return err
}

func (svc *SysDeptSvc) GetDeptList(deptName, deptCode string, pageNo, pageSize int) ([]model.SysDept, int64, error) {
    s := dao.NewStatement()
    if deptName != "" {
        s.Where("dept_name LIKE ?", "%"+deptName+"%")
    }
    if deptCode != "" {
        s.Where("dept_code LIKE ?", "%"+deptCode+"%")
    }
    s.Pagination(pageNo, pageSize)
    s.OrderBy("display_order DESC")
    return dao.QueryList[model.SysDept](s)
}

func (svc *SysDeptSvc) GetDeptById(id any) (model.SysDept, error) {
    return dao.QueryById[model.SysDept](id)
}

func (svc *SysDeptSvc) AddDept(dept dto.AddDeptRequest) error {
    s := dao.NewStatement()
    s.Where("dept_code = ?", dept.DeptCode)
    duplicateData, err := dao.Count[model.SysDept](s)
    if err != nil {
        return err
    }
    if duplicateData > 0 {
        return exception.NewException("部门标识已存在")
    }
    return dao.Create(&dept, "sys_dept").Error
}

func (svc *SysDeptSvc) EditDept(dept dto.EditDeptRequest) error {
    d, err := dao.QueryById[model.SysDept](dept.Id)
    if err != nil {
        return err
    }
    if d.DeptCode == global.SystemAdmin {
        return exception.NewException("系统内置部门，不允许修改")
    }
    return dao.Save(&dept, "sys_dept").Error
}

func (svc *SysDeptSvc) DelDept(id any) error {
    d, err := dao.QueryById[model.SysDept](id)
    if err != nil {
        return err
    }
    if d.DeptCode == global.SystemAdmin {
        return exception.NewException("系统内置部门，不允许删除")
    }
    s := dao.NewStatement()
    s.Where("parent_id = ?", id)
    childDeptCount, err := dao.Count[model.SysDept](s)
    if err != nil {
        return err
    }
    if childDeptCount > 0 {
        return exception.NewException("存在子部门，不允许删除")
    }
    s.Init()
    s.Where("dept_id = ?", id)
    userCount, err := dao.Count[model.SysUser](s)
    if err != nil {
        return err
    }
    if userCount > 0 {
        return exception.NewException("存在用户，不允许删除")
    }
    return dao.DelById[model.SysDept](id)
}
