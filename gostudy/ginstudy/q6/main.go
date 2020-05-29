package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 给表达设置上传大小(默认32MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8Mib

	router.POST("/upload", func(c *gin.Context) {
		/*
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s", file.Filename)
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
			})
		*/
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})

	})
	router.Run(":9090")
}
