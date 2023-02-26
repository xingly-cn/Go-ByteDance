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

	sbrGroup := r.Group("/sbr")
	{
		sbrGroup.GET("/getNum", controller.GetNum)
	}

}
