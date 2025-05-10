package service

import (
	"bosh-admin/core/exception"
	"bosh-admin/dao"
	"bosh-admin/dao/dto"
	"bosh-admin/dao/model"
)

type SysMenuSvc struct{}

func NewSysMenuSvc() *SysMenuSvc {
	return &SysMenuSvc{}
}

// GetMenuTree 获取菜单树
func (svc *SysMenuSvc) GetMenuTree() ([]model.SysMenu, error) {
	treeMap, err := getMenuTreeMap()
	menuTree := treeMap[0]
	for i := 0; i < len(menuTree); i++ {
		err = getMenuChildrenList(&menuTree[i], treeMap)
	}
	return menuTree, err
}

// getMenuTreeMap 获取菜单Map
func getMenuTreeMap() (map[uint][]model.SysMenu, error) {
	var allMenus []model.SysMenu
	treeMap := make(map[uint][]model.SysMenu)
	err := dao.GormDB().Order("display_order DESC").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// getMenuChildrenList 获取子菜单列表
func getMenuChildrenList(menu *model.SysMenu, treeMap map[uint][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = getMenuChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (svc *SysMenuSvc) GetMenus(title string, pageNo, pageSize int) ([]model.SysMenu, int64, error) {
	s := dao.NewStatement()
	if title != "" {
		s.Where("title LIKE ?", "%"+title+"%")
	}
	s.Pagination(pageNo, pageSize)
	s.OrderBy("display_order DESC")
	menus, total, err := dao.QueryList(model.SysMenu{}, s)
	return menus, total, err
}

func (svc *SysMenuSvc) GetMenuById(id any) (model.SysMenu, error) {
	return dao.QueryById[model.SysMenu](id)
}

func (svc *SysMenuSvc) AddMenu(menu model.SysMenu) error {
	s := dao.NewStatement()
	if menu.MenuType < 3 {
		s.Where("a.menu_type < ?", 3)
		s.Where("a.name = ?", menu.Name)
		duplicateName, err := dao.Count(model.SysMenu{}, s)
		if err != nil {
			return err
		}
		if duplicateName > 0 {
			return exception.NewException("路由名称已存在，必须保持唯一")
		}
	} else {
		s.Where("a.parent_id = ?", menu.ParentId)
		s.Where("a.auth_code = ?", menu.AuthCode)
		duplicateAuth, err := dao.Count(model.SysMenu{}, s)
		if err != nil {
			return err
		}
		if duplicateAuth > 0 {
			return exception.NewException("权限标识已存在")
		}
	}
	return dao.Create(&menu).Error
}

func (svc *SysMenuSvc) EditMenu(menu model.SysMenu) error {
	m, err := dao.QueryById[model.SysMenu](menu.Id)
	if err != nil {
		return err
	}
	if m.MenuType < 3 && m.Name != menu.Name {
		s := dao.NewStatement()
		s.Where("a.menu_type < ?", 3)
		s.Where("a.name = ?", menu.Name)
		duplicateName, err := dao.Count(model.SysMenu{}, s)
		if err != nil {
			return err
		}
		if duplicateName > 0 {
			return exception.NewException("路由名称已存在，必须保持唯一")
		}
	}
	return dao.Save(&menu).Error
}

func (svc *SysMenuSvc) DelMenu(id any) error {
	menu, err := dao.QueryById[model.SysMenu](id)
	if err != nil {
		return err
	}
	if menu.MenuType < 3 {
		// 是否有非按钮子菜单
		s := dao.NewStatement()
		s.Where("parent_id = ?", id)
		s.Where("menu_type != ?", 3)
		children, err := dao.Count(model.SysMenu{}, s)
		if err != nil {
			return err
		}
		if children > 0 {
			return exception.NewException("存在子菜单，无法删除")
		}
	}
	tx := dao.Begin()
	// 删除菜单
	err = tx.Delete(&model.SysMenu{}, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// TODO 删除角色-菜单关联
	//err = tx.Where("menu_id = ?", id).Delete(&model.SysRoleMenu{}).Error
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	// 删除按钮子菜单及角色-按钮子菜单关联
	if menu.MenuType < 3 {
		var btnIds []uint
		err = tx.Model(&model.SysMenu{}).Where("parent_id = ?", id).Where("menu_type = ?", 3).Find(&btnIds).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		if len(btnIds) > 0 {
			// 删除按钮子菜单
			err = tx.Delete(&model.SysMenu{}, btnIds).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			// TODO 删除角色-按钮子菜单关联
			//err = tx.Where("menu_id IN ?", btnIds).Delete(&model.SysRoleMenu{}).Error
			//if err != nil {
			//	tx.Rollback()
			//	return err
			//}
		}
	}
	tx.Commit()
	return nil
}

// getAsyncRoutesChildrenList 获取pure admin子菜单列表
func getAsyncRoutesChildrenList(menu *dto.PureMenu, treeMap map[uint][]dto.PureMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getAsyncRoutesChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// TODO 获取pure admin菜单
//func (svc *SysMenuSvc) GetAsyncRoutes(roleId uint, roleCode string) ([]dto.PureMenu, error) {
//
//}
