package model

import "bosh-admin/dao"

// SysUser 系统用户
type SysUser struct {
	dao.BasicModel
	Username      string         `gorm:"username" json:"username"`                    // 账号
	Password      string         `gorm:"password" json:"-"`                           // 密码
	PwdUpdatedAt  dao.CustomTime `gorm:"pwd_updated_at" json:"pwdUpdatedAt"`          // 密码修改时间
	PwdRemainTime int            `gorm:"pwd_remain_time" json:"-"`                    // 剩余尝试次数
	Avatar        string         `gorm:"avatar" json:"avatar"`                        // 头像
	Nickname      string         `gorm:"nickname" json:"nickname"`                    // 昵称
	Gender        int            `gorm:"gender" json:"gender" form:"gender"`          // 性别 0未知 1男 2女
	Birthday      dao.CustomDate `gorm:"birthday" json:"birthday"`                    // 生日
	Email         string         `gorm:"email" json:"email"`                          // 邮箱
	Mobile        string         `gorm:"mobile" json:"mobile"`                        // 联系方式
	Introduce     string         `gorm:"introduce" json:"introduce"`                  // 个人简介
	Status        int            `gorm:"status" json:"status"`                        // 状态 0冻结 1正常
	RoleId        uint           `gorm:"role_id" json:"roleId"`                       // 角色id
	DeptId        uint           `gorm:"dept_id" json:"deptId"`                       // 部门id
	Remark        string         `gorm:"remark" json:"remark"`                        // 备注
	Role          SysRole        `gorm:"foreignKey:RoleId;references:Id" json:"role"` // 角色
	Dept          SysDept        `gorm:"foreignKey:DeptId;references:Id" json:"dept"` // 部门
}

// TableName sys_user
func (SysUser) TableName() string {
	return "sys_user"
}
