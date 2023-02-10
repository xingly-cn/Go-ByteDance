package main

import (
	"ft-security/controller"
	"ft-security/router"
	"github.com/gin-gonic/gin"
)

func main() {

	controller.InitMySQL()

	r := gin.Default()

	router.InitRouter(r)

	r.Run(":518")

}
