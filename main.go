package main

import "github.com/gin-gonic/gin"

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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
