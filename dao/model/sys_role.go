package model

import "bosh-admin/dao"

// SysRole 角色
type SysRole struct {
	dao.BasicModel
	RoleName      string `gorm:"role_name" json:"roleName" `           // 角色名称
	RoleCode      string `gorm:"role_code" json:"roleCode" `           // 角色编码
	Status        int    `gorm:"status" json:"status" `                // 状态 0冻结 1正常
	Remark        string `gorm:"remark" json:"remark" `                // 备注
	DefaultRouter string `json:"default_router" json:"defaultRouter" ` // 默认路由
	DataAuth      int    `gorm:"data_auth" json:"dataAuth" `           // 数据权限 0未配置 1全部数据 2本部门数据 3本部门及以下数据 4本人数据 5自定义数据
}

// TableName sys_role
func (SysRole) TableName() string {
	return "sys_role"
}
