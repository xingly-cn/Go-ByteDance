package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserChangeMac(c *gin.Context) {
	user := proto.User{}
	username := c.Query("username")
	mac := c.Query("mac")

	num, _ := rd.Get("user:" + username).Int()
	if num > 3 {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "换绑次数超过三次, 已记录违规操作", user))
		return
	}

	db.Where("username = ?", username).First(&user)
	if user.MacAddress != mac {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "解绑失败", user))
		return
	}
	db.Model(&proto.User{}).Where("username = ?", username).Update("mac_address", "")

	rd.Incr("user:" + username)

	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "解绑成功", user))
}
