package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		_c := c.Copy()
		go func() {
			log.Println("正在处理请求", _c.Request.URL.Path)
			time.Sleep(time.Second * 5)
			log.Println("请求已处理", _c.Request.URL.Path)
		}()
		c.String(http.StatusOK, "done")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
