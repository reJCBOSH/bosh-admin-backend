package service

import (
	"bosh-admin/core/exception"
	"bosh-admin/dao"
	"bosh-admin/dao/dto"
	"bosh-admin/dao/model"
	"bosh-admin/utils"
	"fmt"
)

type SysUserSvc struct{}

func NewSysUserSvc() *SysUserSvc {
	return &SysUserSvc{}
}

func (svc *SysUserSvc) GetUserList(username, nickname string, gender, status *int, roleId, deptId *uint, pageNo, pageSize int) ([]model.SysUser, int64, error) {
	s := dao.NewStatement()
	if username != "" {
		s.Where("username LIKE ?", "%"+username+"%")
	}
	if nickname != "" {
		s.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	if gender != nil {
		s.Where("gender = ?", *gender)
	}
	if status != nil {
		s.Where("status = ?", *status)
	}
	if roleId != nil {
		s.Where("role_id = ?", *roleId)
	}
	if deptId != nil {
		s.Where("dept_id = ?", *deptId)
	}
	s.Pagination(pageNo, pageSize)
	s.Preload("Role")
	s.Preload("Dept")
	return dao.QueryList(model.SysUser{}, s)
}

func (svc *SysUserSvc) GetUserById(id any) (model.SysUser, error) {
	s := dao.NewStatement()
	s.Where("id = ?", id)
	s.Preload("Role")
	s.Preload("Dept")
	return dao.QueryOne(model.SysUser{}, s)
}

func (svc *SysUserSvc) AddUser(user dto.AddUserRequest) error {
	s := dao.NewStatement()
	s.Where("username = ?", user.Username)
	count, err := dao.Count(model.SysUser{}, s)
	if err != nil {
		return err
	}
	if count > 0 {
		return exception.NewException("用户名已存在")
	}
	return dao.Create(user, "sys_user").Error
}

func (svc *SysUserSvc) EditUser(user dto.EditUserRequest) error {
	u, err := dao.QueryById[model.SysUser](user.Id)
	if err != nil {
		return err
	}
	if user.Username != u.Username {
		s := dao.NewStatement()
		s.Where("username = ?", user.Username)
		count, err := dao.Count(model.SysUser{}, s)
		if err != nil {
			return err
		}
		if count > 0 {
			return exception.NewException("用户名已存在")
		}
	}
	return dao.Save(user, "sys_user").Error
}

func (svc *SysUserSvc) DelUser(id any) error {
	// TODO 判断是否是超级管理员或者自杀行为
	return dao.DelById(id)
}

func (svc *SysUserSvc) Login(username, password, captcha, captchaId string) (*model.SysUser, error) {
	if !utils.VerifyCaptcha(captchaId, captcha) {
		return nil, exception.NewException("验证码错误")
	}
	s := dao.NewStatement()
	s.Where("username = ?", username)
	s.Preload("Role")
	s.Preload("Dept")
	user, err := dao.QueryOne(model.SysUser{}, s)
	if err != nil {
		if err == dao.NotFound {
			return nil, exception.NewException("账号或密码错误")
		}
		return nil, err
	}
	if user.Status == 0 {
		return nil, exception.NewException("账号已被冻结, 请联系管理员")
	}
	if ok := utils.BcryptCheck(password, user.Password); !ok {
		if user.PwdRemainTime == 1 {
			user.PwdRemainTime = 0
			user.Status = 0
			if err = dao.Save(user).Error; err != nil {
				return nil, err
			}
			return nil, exception.NewException("账号已被冻结, 请联系管理员")
		}
		if err = dao.GormDB().Model(model.SysUser{}).Where("id = ?", user.Id).UpdateColumn("pwd_remain_time", dao.Expr("pwd_remain_time - ?", 1)).Error; err != nil {
			return nil, err
		}
		return nil, exception.NewException(fmt.Sprintf("密码错误，剩余%d次尝试机会，超出则冻结账号", user.PwdRemainTime-1))
	}
	// TODO JWT

	tx := dao.Begin()
	if user.PwdRemainTime < 5 {
		if err = tx.Model(model.SysUser{}).Where("id = ?", user.Id).UpdateColumn("pwd_remain_time", 5).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	// TODO 增加登录记录
	tx.Commit()
	return &user, nil
}
