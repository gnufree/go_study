package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/moreXML", func(c *gin.Context) {
		// 第二种json
		type MessageRecode struct {
			Name    string
			Message string
			Number  int
		}
		var msg MessageRecode
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		c.XML(http.StatusOK, msg)
	})

	router.Run(":9090")
}
