package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("go-bytedance:go-bytedance@tcp(175.27.243.243:3306)/go-bytedance?charset=utf8"), &gorm.Config{
		SkipDefaultTransaction: true, // 关闭默认事务
		PrepareStmt:            true, // 编译预编译
	})
	if err != nil {
		panic("connect failed")
	}

	// 创建一条
	//temp := &Product{Code: "你好啊", Price: 100}
	//res := db.Create(temp)
	//fmt.Println(res.Error)

	// 查询一条
	u := &Product{}
	db.Where("price > ?", 0).First(&u)
	fmt.Println(*u)

	// 更新数据
	//db.Model(&Product{}).Where("price > ?", 0).Update("code", "我修改了")

	// 物理删除
	//db.Delete(&Product{}, 10) // 删除id = 10

	// 事务
	//db.Transaction(func(tx *gorm.DB) error {
	//	if err := tx.Create(&Product{ID: 111}).Error; err != nil {
	//		return err
	//	}
	//	return nil
	//})

}
