package routers

import (
	"pc_vue/controllers/userinfo"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		// 用户登陆
		userRouter.POST("/login", userinfo.Userinfocontrollers.Login)
	}

}
