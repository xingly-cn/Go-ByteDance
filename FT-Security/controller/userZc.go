package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserZc(c *gin.Context) {
	log.Println("ok")

	username := c.Query("username")
	password := c.Query("password")
	mac := c.Query("mac")

	user := proto.User{}
	db.Where("username = ?", username).First(&user)
	if user.Username != "" {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "账号存在", nil))
		RecordLog(username, "发起注册-账号存在", mac)
		return
	}

	user.Status = true
	user.MacAddress = mac
	user.Username = username
	user.Password = password
	db.Create(&user)

	if user.ID == 0 {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "注册失败", nil))
		RecordLog(username, "发起注册-注册失败", mac)
		return
	}

	RecordLog(username, "发起注册-成功", mac)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "注册成功", user))
}
