package controller

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var rd *redis.Client
var err error

func InitMySQL() {
	db, err = gorm.Open(mysql.Open("root:XNXxnx520@@tcp(gz-cynosdbmysql-grp-lbda0189.sql.tencentcdb.com:27351)/sbr?charset=utf8&parseTime=true"), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 编译预编译
	})
	if err != nil {
		log.Fatal(err)
	}
}

func InitRedis() {
	rd = redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 1})
}
