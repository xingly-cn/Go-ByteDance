package main

import (
	"ft-security/controller"
	"ft-security/router"
	"github.com/gin-gonic/gin"
)

func main() {

	controller.InitMySQL()
	controller.InitRedis()
	controller.AdminIntoLocalMap()

	r := gin.Default()

	router.InitRouter(r)

	r.Run(":518")

}
