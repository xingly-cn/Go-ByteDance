package controller

import (
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecordActUser(c *gin.Context) {
	k := c.Query("k")
	v := c.Query("v")
	activityId := c.Query("activityId")
	log := RecordActivityAndUserLog(k, v, activityId)
	RecordLog(k, "活动与用户绑定", "@all")
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "活动与用户绑定->"+log, nil))
}

func RecordActUserCheck(c *gin.Context) {
	phone := c.Query("k")
	activityId := c.Query("activityId")
	RecordLog(phone, "活动与用户鉴权", "@all")
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "活动与用户资格鉴权", RecordActivityAndUserLogCheck(phone, activityId)))
}
