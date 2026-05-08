package handlers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartServer(serverPort string) {

	gin.SetMode(gin.ReleaseMode)

	rateLimiter := NewRateLimiter(parseRateLimit(os.Getenv("RATE_LIMIT")), time.Minute)

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

	router.Use(rateLimiter.Middleware())

	router.POST("/slink", addLink)
	router.GET("/s/:hash", goToLink)
	router.GET("/gst/:hash", getLinkStats)

	fmt.Println("Server started")
	router.Run(serverPort)
}

func parseRateLimit(s string) int {
	if s == "" {
		return 100
	}
	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 {
		return 100
	}
	return n
}
