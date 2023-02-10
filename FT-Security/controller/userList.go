package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	var users []proto.User
	db.Find(&users)
	if len(users) == 0 {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "无数据", nil))
		return
	}
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "用户列表", users))
}
