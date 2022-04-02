package models

// binding:"required"

type LoginParams struct {
	Account  string `json:"account" form:"account" binding:"required"`     //账号
	PossWord string `json:"password"  form:"password"  binding:"required"` //密码
}
