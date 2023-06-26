package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 设置生产模式，默认 debug 模式
	gin.SetMode(gin.ReleaseMode)

	s := gin.Default()

	s.GET("/hello", func(c *gin.Context) {
		//_, err := c.Writer.WriteString("Hello, Welcome to gin-gonic")
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}
		c.JSON(http.StatusOK, gin.H{
			"name":    "zhangsan",
			"age":     41,
			"hobbies": []string{"jogging", "reading"},
		})
	})

	err := s.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
