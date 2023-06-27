package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name string `uri:"name"`
	Age  int    `uri:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/users/:name", func(c *gin.Context) {
		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			log.Println(err)
		}
		c.JSONP(http.StatusOK, &p)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
