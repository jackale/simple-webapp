package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackale/simple-webapp/src/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("src/view/*")

	router.GET("/", controller.IndexGET)
	router.Run(":5555")
}
