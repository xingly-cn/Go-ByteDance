package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open("go-bytedance:go-bytedance@tcp(175.27.243.243:3306)/go-bytedance?charset=utf8"), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 编译预编译
	})
	return db, err
}
