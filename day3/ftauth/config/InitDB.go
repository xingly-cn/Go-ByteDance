package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open("root:XNXxnx520@@tcp(gz-cynosdbmysql-grp-lbda0189.sql.tencentcdb.com:27351)/ft-ai?charset=utf8"), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 编译预编译
	})
	return db, err
}
