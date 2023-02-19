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
	for i := range users {
		users[i].Password = "********"
	}
	RecordLog("管理员", "查看用户列表", "@all")

	// todo 临时更新管理员内存表
	AdminIntoLocalMap()

	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "用户列表", users))
}
