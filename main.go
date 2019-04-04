package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("call main()")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:9090
	r.Run(":9090")
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// TODO

	return r
}
