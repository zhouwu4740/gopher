package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//r := InjectLoggerEngine()
	r := gin.Default()
	r.Use(Logger())
	r.GET("/hello", func(c *gin.Context) {
		spanId, _ := c.Get("spanId")
		log.Println("spanId = ", spanId)
		c.String(http.StatusOK, "OK")
	})
	r.Run(":8080")
}
