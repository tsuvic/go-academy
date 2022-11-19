package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/echo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Request GET",
		})
	})

	r.GET("/book", func(c *gin.Context) {
		title := c.DefaultQuery("title", "None")
		bookid := c.Query("bookid")
		c.String(http.StatusOK, "Request %s %s", bookid, title)
	})

	r.Run()
}
