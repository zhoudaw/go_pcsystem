package models

// binding:"required"

type LoginParams struct {
	Account  string `json:"account" form:"account" binding:"required"`     //账号
	PossWord string `json:"password"  form:"password"  binding:"required"` //密码
}

// 表示配置操作数据库的表名称
func (LoginParams) TableName() string {
	return "user_info"
}
