package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestPayload represents the expected JSON payload for the POST handler
/*
The struct you provided is a Go struct with struct tags (annotations) that are used for JSON serialization and data validation in the Gin framework
type RegisterRequest struct {
    Name      string `json:"name" binding:"required,min=3,max=20"`  // Must be 3-20 characters
    Email     string `json:"email" binding:"required,email"`        // Must be a valid email
    Password  string `json:"password" binding:"required,min=6"`     // Minimum 6 characters
    Age       int    `json:"age" binding:"gte=18,lte=100"`          // Age must be between 18-100
    IsAdmin   bool   `json:"is_admin"`                              // Optional field
}
*/
type RequestPayload struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// helloHandler handles GET requests to the /api/v0/gin endpoint
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Gin web framework",
	})
}

// postHandler handles POST requests to the /api/v0/gin endpoint
func postHandler(c *gin.Context) {
	var payload RequestPayload

	// Bind JSON payload to the struct and validate it
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request payload",
			"details": err.Error(),
		})
		return
	}

	// Respond with the received data
	c.JSON(http.StatusOK, gin.H{
		"message": "Payload received successfully",
		"data":    payload,
	})
}

func main() {
	// Create a new Gin router with default middleware (logging and recovery)
	r := gin.Default()

	// Group API routes for better organization
	api := r.Group("/api/v0/gin")
	{
		api.GET("", helloHandler)
		api.POST("", postHandler)
	}

	// Start the server and handle errors
	if err := r.Run(); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
