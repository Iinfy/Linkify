package handlers

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartServer(serverPort string) {
	// Удалите cors.Default() — используем только одну настройку

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("CORS_HOST"),
		},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(func(c *gin.Context) {
		c.Header("Referrer-Policy", "no-referrer")
		c.Next()
	})

	router.POST("/slink", addLink)
	router.GET("/s/:hash", goToLink)
	router.GET("/short/:hash", getLinkStats)
	// router.GET("/qr", createQR)

	fmt.Println("Server started")
	router.Run(serverPort)
}
