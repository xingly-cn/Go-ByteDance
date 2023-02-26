package main

import (
	"ftGame/controller"
	"ftGame/router"
	"github.com/gin-gonic/gin"
)

func main() {

	controller.InitMySQL()
	controller.InitRedis()

	r := gin.Default()

	router.InitRouter(r)

	r.Run(":888")

}
