package main

import (
	"pc_vue/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 路由配置
	//用户
	routers.UserRouters(r)

	r.Run(":8000")
}
