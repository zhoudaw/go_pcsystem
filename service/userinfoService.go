package service

import (
	"pc_vue/dao/admininfo"
	"pc_vue/models"
)

// 获取账号密码
func AdminLogin(loginParmas *models.LoginParams) {
	admininfo.GetAdmin(loginParmas)
}
