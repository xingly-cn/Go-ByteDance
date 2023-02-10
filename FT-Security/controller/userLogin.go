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

	if user.MacAddress == "" {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "登录成功", user))
		db.Model(&proto.User{}).Where("username = ?", username).Update("mac_address", mac)
		return
	}

	if user.Username == "" || user.Status == false || user.UseDay <= 0 || (user.Password != password) || (user.MacAddress != mac) {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "账号冻结", nil))
		return
	}

	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "登录成功", user))
}
