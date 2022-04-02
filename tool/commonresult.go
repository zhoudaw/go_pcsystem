package tool

import (
	"net/http"
	"pc_vue/models"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 200
	FAILED  = 500
)

func Success(c *gin.Context, data interface{}) {
	result := &models.Result{
		Code: SUCCESS,
		Msg:  "成功",
		Data: data,
	}
	c.JSON(http.StatusOK, result)

}
func Failed(c *gin.Context, data interface{}) {
	result := &models.Result{
		Code: FAILED,
		Msg:  "失败",
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}
