package main

import (
	"fmt"
	"net/http"
	"url-shortner/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting URL shortener Service")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin-based URL Shortner Service",
		})
	})

	router.POST("/shorten", handlers.ShortenURL)

	router.Run(":8080")
}
