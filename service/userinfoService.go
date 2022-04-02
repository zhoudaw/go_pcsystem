package service

import (
	"pc_vue/dao/userinfodao"
	"pc_vue/models"
)

// 获取账号密码
func AdminLogin(loginParmas *models.LoginParams) {
	userinfodao.GetAdmin(loginParmas)
}
