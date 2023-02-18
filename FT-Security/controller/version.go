package controller

import (
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersion(c *gin.Context) {
	v, _ := rd.Get("version").Result()
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "当前版本", v))
}
