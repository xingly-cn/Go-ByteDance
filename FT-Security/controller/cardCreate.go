package controller

import (
	"ft-security/proto"
	"ft-security/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateCard(c *gin.Context) {
	num, _ := strconv.Atoi(c.Query("num"))
	score, _ := strconv.Atoi(c.Query("score"))
	var list []string
	for i := 0; i < num; i++ {
		uid := utils.GetUuid()
		list = append(list, uid)
		t := &proto.Card{Uid: uid, Score: int64(score), Status: false, UseTime: time.Now()}
		db.Create(&t)
	}
	RecordLog("管理员", "生成卡密->"+strconv.Itoa(num)+"/"+strconv.Itoa(num), "@all")
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "生成成功", list))
}
