package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1", func(c *gin.Context) {
		c.String(http.StatusOK, "group v1\n")
	})
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "hello group v1")
		})
		v1.GET("/hi", func(c *gin.Context) {
			c.String(http.StatusOK, "hi group v1")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "hello group v2")
		})
		v2.GET("/hi", func(c *gin.Context) {
			c.String(http.StatusOK, "hi group v2")
		})
	}

	r.Run(":8080")
}
