package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

func HealthCheck(c *gin.Context) {
	c.String(200, "ok")
}

func main() {
	r := gin.Default()

	r.GET("/", Hello)
	r.GET("/health_check", HealthCheck)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
