package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", login)
		v1.POST("/submit", submit)
		v1.POST("/read", read)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
		v2.POST("/read", read)
	}

	router.Run(":9090")
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login succ",
	})
}

func submit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit succ",
	})
}

func read(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit succ",
	})
}
