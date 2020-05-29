package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get user info",
		})
	})
	r.POST("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "create user info",
		})
	})
	r.PUT("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update user info",
		})
	})
	r.DELETE("/user/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete user info",
		})
	})
	r.Run(":9090")
}
