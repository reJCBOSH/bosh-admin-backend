package utils

import "bosh-admin/global"

// IsProd 是否生产环境
func IsProd() bool {
	if global.Config.Server.Env == global.PROD {
		return true
	}
	return false
}
