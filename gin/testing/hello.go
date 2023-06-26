package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func setupRouter() *gin.Engine {
	s := gin.Default()

	s.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello gin-gonic")
	})

	return s
}

func main() {
	s := setupRouter()

	err := s.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
