package main

import (
	"day3/ftai/config"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type User struct {
	ID          int    `json:"ID"`
	U           string `json:"u"`
	P           string `json:"p"`
	Flag        bool   `json:"flag"`
	Time        int    `json:"time"`
	LastLoginIP string `json:"last_login_ip"`
}

type Card struct {
	ID      int    `json:"ID"`
	Cdk     string `json:"cdk"`
	UseTime string `json:"use_time"`
	UseID   string `json:"use_id"`
	Score   int    `json:"score"`
}

var (
	typer    string
	username string
	password string
	c        string
	ip       string
	// 缓存
	user User
	card Card
)

func main() {
	// 启动页
	log.Println("方糖云授权工具 Version：1.0.0 - 卡密地址：card.asugar.cn")

	// link to mysql
	db, err := config.Start()
	if err != nil {
		log.Fatal("Error which link to MYSQL...")
	}
	log.Println("Success which link to MYSQL...")
	log.Println("输入命令 -> 注册账号：1 卡密充值：2 查询天数：3 退出：4")

	// type select
	fmt.Scanf("%s\n", &typer)
	for {
		switch typer {
		case "1":
			log.Println("请输入账号")
			fmt.Scanf("%s\n", &username)
			log.Println("请输入密码")
			fmt.Scanf("%s\n", &password)
			flag := checkUserNameIsOk(username, password, db)
			if flag {
				log.Println("注册成功")
			} else {
				log.Println("用户名已存在")
			}
			log.Println("----------------------------------------")
			log.Println("输入命令 -> 注册账号：1 卡密充值：2 查询天数：3 退出：4")
			fmt.Scanf("%s\n", &typer)
		case "2":
			log.Println("请输入账号")
			fmt.Scanf("%s\n", &username)
			log.Println("请输入卡密")
			fmt.Scanf("%s\n", &c)
			flag, timer := payUsername(username, c, db)
			if flag {
				log.Println("充值成功 -> 天数+" + strconv.Itoa(timer))
			} else {
				log.Println("充值失败, 卡密错误或已使用")
			}
			log.Println("----------------------------------------")
			log.Println("输入命令 -> 注册账号：1 卡密充值：2 查询天数：3 退出：4")
			fmt.Scanf("%s\n", &typer)
		case "3":
			log.Println("请输入账号")
			fmt.Scanf("%s\n", &username)
			log.Println("请输入密码")
			fmt.Scanf("%s\n", &password)
			flag := login(username, password, db)
			if flag {
				log.Println("账号状态：", user.Flag, "剩余天数：", user.Time, "上次登录IP：", user.LastLoginIP)
			} else {
				log.Println("登录失败")
			}
			log.Println("----------------------------------------")
			log.Println("输入命令 -> 注册账号：1 卡密充值：2 查询天数：3 退出：4")
			fmt.Scanf("%s\n", &typer)
		case "4":
			log.Println("欢迎使用, 如有问题 -> Email：admin@asugar.cn, 投Bug有奖励~")
			return
		}
	}

}

// 注册
func checkUserNameIsOk(u string, p string, db *gorm.DB) bool {
	db.Where("u = ?", u).First(&user)
	if user.ID == 0 {
		db.Create(&User{U: u, P: p, Flag: true, Time: 0})
		return true
	}
	return false
}

// 充值
func payUsername(u string, c string, db *gorm.DB) (bool, int) {
	db.Where("cdk = ?", c).First(&card)
	if card.ID == 0 || card.UseID != "" {
		return false, 0
	}
	score := card.Score
	db.Where("u = ?", u).First(&user)
	db.Model(&User{}).Where("u = ?", u).Update("time", user.Time+score)
	db.Model(&Card{}).Where("cdk = ?", c).Updates(Card{UseID: u, UseTime: time.Now().String()})
	return true, score
}

// 登录
func login(u string, p string, db *gorm.DB) bool {
	db.Where("u = ? and p = ?", u, p).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
