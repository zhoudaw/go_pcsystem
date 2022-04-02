package userinfo

import (
	"pc_vue/models"
	"pc_vue/service"
	"pc_vue/tool"

	"github.com/gin-gonic/gin"
)

var Userinfocontrollers = &UserInfoControllers{}

type UserInfoControllers struct{}

//用户登陆
func (userinfo *UserInfoControllers) Login(c *gin.Context) {
	loginParmas := &models.LoginParams{}
	err := c.ShouldBindJSON(&loginParmas)
	if err != nil {
		tool.Failed(c, err)
		return
	}
	//如果有接收到参数执行函数
	service.AdminLogin(loginParmas)
	tool.Success(c, loginParmas)
}
