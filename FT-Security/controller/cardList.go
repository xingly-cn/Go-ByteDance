package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCardList(c *gin.Context) {
	var cards []proto.Card
	db.Order("status").Find(&cards)
	if len(cards) == 0 {
		c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "无数据", nil))
		return
	}
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "卡密列表", cards))
}
