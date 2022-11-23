package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(EchoMiddleware())

	r.GET("/echo", func(c *gin.Context) {
		fmt.Println("exec domain func")
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})

	r.Run()
}

func EchoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Method, c.Request.Host, c.Request.URL.Path)

		fmt.Println("exec common func before domain func")

		c.Next()

		fmt.Println("exec common func after domain func")
	}
}
