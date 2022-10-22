package main

import (
	"gin/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		hello := new(controllers.HelloWorldController)
		v1.GET("/hello", hello.Default)

		user := new(controllers.UserController)
		v1.POST("/signup", user.Signup)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	router.Run(":5001")
}
