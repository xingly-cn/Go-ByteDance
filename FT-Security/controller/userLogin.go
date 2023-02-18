package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	mac := c.Query("mac")

	user := proto.User{}
	db.Where("username = ?", username).First(&user)

	if user.Status == false || user.Username == "" || user.UseDay <= 0 || (user.Password != password) {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "账号冻结", nil))
		RecordLog(username, "发起登录-账号冻结", mac)
		return
	}

	if user.MacAddress == "" {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "登录成功", user))
		db.Model(&proto.User{}).Where("username = ?", username).Update("mac_address", mac)
		RecordLog(username, "发起登录-成功", mac)
		return
	}

	if user.MacAddress != mac {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "账号冻结", nil))
		RecordLog(username, "发起登录-账号冻结", mac)
		return
	}
	RecordLog(username, "发起登录-成功", mac)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "登录成功", user))
}
