package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		message := c.PostForm("message")
		title := c.DefaultPostForm("title", "none")
		c.String(http.StatusOK, "Request %s %s", message, title)
	})

	r.Run()
}
