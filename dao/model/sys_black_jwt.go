package model

import "bosh-admin/dao"

type SysBlackJwt struct {
    dao.BasicModel
    BlackJwt  string `gorm:"comment:作废JWT" json:"blackJwt"`
    BlackUUID string `gorm:"black_uuid;comment:作废UUID" json:"blackUUID"`
}
