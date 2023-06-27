package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `json:"name" form:"name"`
	Age      int       `json:"age" form:"age"`
	Gender   string    `json:"gender" form:"gender"`
	Birthday time.Time `json:"birthday" form:"birthday" time_format:"2006-01-02"`
	Hobbies  []string  `json:"hobbies" form:"hobbies"` // 复选框
}

type Student struct {
	*Person
	Id    int    `json:"id" form:"id"`
	Score string `json:"score" form:"score"`
}

func (s *Student) String() string {
	return fmt.Sprintf("%d %s %d %s %s %s %s", s.Id, s.Name, s.Age, s.Birthday, s.Gender, s.Score, s.Hobbies)
}

func main() {
	r := gin.Default()

	r.POST("/student", func(c *gin.Context) {
		var s Student
		if err := c.Bind(&s); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, s)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
