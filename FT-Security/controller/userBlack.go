package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BlackUser(c *gin.Context) {
	user := proto.User{}
	username := c.Query("username")
	db.Where("username = ?", username).First(&user)
	db.Model(&proto.User{}).Where("username = ?", username).Update("status", false)
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "加黑成功", user))
}
