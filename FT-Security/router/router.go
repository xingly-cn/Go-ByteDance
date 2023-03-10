package router

import (
	"ft-security/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.Static("/log", "./log")

	userGroup := r.Group("/user")
	{
		userGroup.GET("/zc", controller.UserZc)                      // 注册
		userGroup.GET("/login", controller.UserLogin)                // 登录
		userGroup.GET("/changeMacAddress", controller.UserChangeMac) // 换绑
		userGroup.GET("/list", controller.GetUserList)               // 管理员-用户列表
		userGroup.GET("/black", controller.BlackUser)                // 黑名单
		userGroup.GET("/times", controller.UserPayDays)              // 管理员-批量扣除
		userGroup.GET("/record", controller.RecordActUser)           // 管理员-批量扣除
		userGroup.GET("/check", controller.RecordActUserCheck)       // 管理员-批量扣除
	}

	cardGroup := r.Group("/card")
	{
		cardGroup.GET("/use", controller.CardUse)       // 使用卡密
		cardGroup.GET("/create", controller.CreateCard) // 生成卡密
		cardGroup.GET("/list", controller.GetCardList)  // 卡密列表
	}

	versionGroup := r.Group("version")
	{
		versionGroup.GET("update", controller.GetVersion)
	}
}
