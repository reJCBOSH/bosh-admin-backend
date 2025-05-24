package model

import "bosh-admin/dao"

type SysBlackJwt struct {
	dao.BasicModel
	BlackJwt  string `gorm:"black_jwt" json:"blackJwt"`   // 作废JWT
	BlackUUID string `gorm:"black_uuid" json:"blackUuid"` // 作废UUID
}

func (SysBlackJwt) TableName() string {
	return "sys_black_jwt"
}
