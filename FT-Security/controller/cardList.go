package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCardList(c *gin.Context) {
	var cards []proto.Card
	var user proto.User
	un := c.Query("username")
	db.Where("username = ?", un).First(&user)

	if user.UseDay < 500 {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "非管理员, 违规操作已被记录", nil))
		return
	}

	db.Order("status").Find(&cards)
	RecordLog("管理员", "查询卡密列表", "@all")
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "卡密列表", cards))
}
