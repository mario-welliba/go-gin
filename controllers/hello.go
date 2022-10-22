package controllers

import (
	// Import the Gin library
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorldController will hold the methods to the
type HelloWorldController struct{}

// Default controller handles returning the hello world JSON response
func (h *HelloWorldController) Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world, climate change is real"})
}
