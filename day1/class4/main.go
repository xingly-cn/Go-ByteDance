package main

import (
	"day1/class4/entity"
	"day1/class4/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {

	if err := entity.Init("./data/"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := gin.Default()
	r.GET("/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		id, _ := strconv.Atoi(topicId)
		data, _ := service.QueryPageInfo(int64(id))
		c.JSON(200, data)
	})

	r.Run()

}
