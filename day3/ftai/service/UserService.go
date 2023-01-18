package service

import (
	"day3/ftai/entity"
	"encoding/base64"
	"gorm.io/gorm"
)

// RegisterOrLogin 注册Or登录
func RegisterOrLogin(username string, password string, db *gorm.DB) (tokenStr string, err string) {
	var user entity.User
	db.Where("username = ?", username).Limit(1).Find(&user)
	if user.ID == 0 {
		// 注册
		db.Create(&entity.User{Username: username, Password: password, Flag: true})
		// 混淆登录
		token := base64.URLEncoding.EncodeToString([]byte(username + "-" + password))
		token = "f" + token
		return token, "true"
	}
	// 账号密码校验
	if password != user.Password {
		return "账号或密码错误", "false"
	}

	// 混淆登录
	token := base64.URLEncoding.EncodeToString([]byte(user.Username))
	token = "f" + token
	return token, "true"
}

// UserPay 充值
func UserPay(name string, card string, db *gorm.DB) (msg string) {
	// 卡密检查
	if card[0] != 'F' && card[0] != 'T' {
		return "系统异常"
	}
	// 账号查询
	var user entity.User
	db.Where("username = ?", name).First(&user)
	if user.ID == 0 {
		return "用户不存在"
	}
	// 充值
	db.Model(&entity.User{}).Where("username = ?", name).Update("balance", user.Balance+1000)
	return "充值成功"
}

// Token2User 转用户
func Token2User(token string, user *entity.User, db *gorm.DB) {
	token = token[1:]
	tokenDecode, _ := base64.URLEncoding.DecodeString(token)
	username := string(tokenDecode)
	db.Where("username = ?", username).First(&user)
	user.Password = "******"
}
