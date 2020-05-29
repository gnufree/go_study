package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someJSON", func(c *gin.Context) {
		// 第一种方式，自己拼json
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK})
	})
	router.GET("/moreJSON", func(c *gin.Context) {
		// 第二种json
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	router.Run(":9090")
}
