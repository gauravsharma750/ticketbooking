package main

import (
	"fmt"

	"github.com/gauravshrma1988/ticketbooking/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig() // Load configuration settings
	// Print the loaded configuration for debugging purposes
	fmt.Printf("Loaded config: %+v\n", cfg)

	r := gin.Default() /// Initialize a new Gin router
	r.GET("/health", func(c *gin.Context) {
		// Define a health check endpoint
		// The endpoint responds with a 200 OK status code and a JSON object.
		// The JSON object contains a single key "status" with the value "ok".
		// This indicates that the service is healthy and operational.
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8080") // Start the Gin server on the default port (8080)
}
