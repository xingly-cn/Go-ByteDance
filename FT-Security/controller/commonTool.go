package controller

import (
	"ft-security/proto"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB
var rd *redis.Client
var err error
var adminPhoneMap map[string]string

func InitMySQL() {
	db, err = gorm.Open(mysql.Open("root:XNXxnx520@@tcp(gz-cynosdbmysql-grp-lbda0189.sql.tencentcdb.com:27351)/ft-security?charset=utf8&parseTime=true"), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 编译预编译
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(proto.Dict{})
}

func InitRedis() {
	rd = redis.NewClient(&redis.Options{Addr: "175.27.243.243:6379", Password: "213879", DB: 1})
}

func RecordLog(uid string, msg string, mac string) {
	log := proto.Log{
		Uid:        uid,
		Msg:        msg,
		MacAddress: mac,
		UseTime:    time.Now(),
	}
	db.Create(&log)
}

func AdminIntoLocalMap() {
	if adminPhoneMap == nil {
		adminPhoneMap = make(map[string]string)
	}

	var adminUser []proto.AdminUser
	db.Find(&adminUser)
	log.Println(adminUser)
	for _, item := range adminUser {
		adminPhoneMap[item.Phone] = "access"
	}
}

func RecordActivityAndUserLog(k string, v string, activityId string) string {
	dict := proto.Dict{
		K:          k,
		V:          v,
		ActivityId: activityId,
		UseTime:    time.Now(),
	}
	db.Create(&dict)
	err := rd.Incr(k) // 增加个人次数
	return err.String()
}

func RecordActivityAndUserLogCheck(phone string, activityId string) bool {
	log.Println("传入参数", phone, activityId)
	var activity proto.Activity
	db.Where("id = ?", activityId).First(&activity)
	dictNum, _ := rd.Get(phone).Int()

	if dictNum == 0 {
		return true
	} else {
		if adminPhoneMap[phone] == "access" {
			log.Println(phone, "通过管理员测试", dictNum, activity.Num)
			// 判断超出
			if dictNum < activity.Num {
				return true
			}
		}
	}
	return false
}
