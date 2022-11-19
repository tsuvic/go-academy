package main

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ginを使ってみました。",
		})
	})
	r.Run()
}
