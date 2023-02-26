package controller

import (
	"ftGame/proto"
	"ftGame/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNum(c *gin.Context) {
	var num int64
	db.Model(proto.User{}).Count(&num)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "查询成功", num))
}
