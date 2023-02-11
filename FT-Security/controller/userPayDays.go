package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserPayDays(c *gin.Context) {
	var users []proto.User
	db.Find(&users)
	if len(users) == 0 {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "无数据", nil))
		return
	}
	for _, item := range users {
		t := item.UseDay - 1
		if t < 0 {
			continue
		}
		db.Model(&proto.User{}).Where("username = ?", item.Username).Update("use_day", t)
	}
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "批量扣除", nil))
}
