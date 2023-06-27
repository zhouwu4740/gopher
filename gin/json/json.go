package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		u := gin.H{
			"1": "张三",
			"2": "李四",
			"3": "王五",
		}
		//c.AsciiJSON(http.StatusOK, &u)
		//c.JSON(http.StatusOK, &u)
		//c.JSONP(http.StatusOK, &u)
		c.SecureJSON(http.StatusOK, &u)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
