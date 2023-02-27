package router

import (
	"ftGame/controller"
	"ftGame/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {

	// r.Static("/log", "./log")

	// 控制中心
	centerGroup := r.Group("/center")
	{
		centerGroup.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "数据湖中心运行正常", nil))
		})
	}
	// 苏泊尔
	sbrGroup := r.Group("/sbr")
	{
		sbrGroup.GET("/getNum", controller.GetNum)
	}
	// 广汽传祺
	gqGroup := r.Group("/gqcq")
	{
		gqGroup.GET("/getList", controller.GetAccountList)
	}
	// e网通
	eGroup := r.Group("/e")
	{
		eGroup.GET("/getBill", controller.GetBill)
		eGroup.GET("/getMoney", controller.GetMoney)
		eGroup.GET("/Loss", controller.Loss)
		eGroup.GET("/getCard", controller.GetCard)
		eGroup.GET("/getUser", controller.GetUser)
	}
}
