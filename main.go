package main

import (
	"github.com/gin-gonic/gin"
	"github.com/timdevir07/Event-Booking-System-/config"
)

func main() {
	config.LoadEnv()
	config.ConnectMongo()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "MongoDB Connected — Event Booking Running")
	})
	r.Run(":8080")
}
