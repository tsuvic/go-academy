package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "./task11/"+"output.txt")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", "output.txt"))

	})

	r.Run()
}
