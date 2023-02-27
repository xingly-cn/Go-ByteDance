package controller

import (
	"fmt"
	"ftGame/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetAccountList(c *gin.Context) {
	list, _ := rd.SMembers("广汽传祺").Result()
	var res []string
	for _, item := range list {
		t := strings.Split(item, "-")[1]
		res = append(res, t)
		fmt.Println(t)
	}
	c.JSON(http.StatusOK, utils.Tell(http.StatusOK, "账号列表", res))
}
