package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CardUse(c *gin.Context) {
	var user proto.User
	var card proto.Card

	username := c.Query("username")
	uid := c.Query("card")

	// pre check
	db.Where("username = ?", username).First(&user)
	if user.Username == "" {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "用户异常", nil))
		return
	}
	db.Where("uid = ?", uid).First(&card)
	if card.Uid == "" || card.Status == true || card.UserId != "" {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "卡密异常", nil))
		return
	}

	// do task
	card.UserId = username
	card.Status = true
	card.UseTime = time.Now()
	db.Model(&proto.Card{}).Where("uid = ?", uid).Updates(card) //todo
	user.UseDay += card.Score
	db.Model(&proto.User{}).Where("username = ?", username).Updates(user)

	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "充值成功", user))
}
