package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//全局调用 *gorm.DB
var DB *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:12345678@tcp(127.0.0.1:3306)/storehouse?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
