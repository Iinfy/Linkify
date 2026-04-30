package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartServer(serverPort string) {
	router.POST("/slink", addLink)

	router.GET("/s/:hash", goToLink)

	fmt.Println("Server started")
	router.Run(":8080")
}
