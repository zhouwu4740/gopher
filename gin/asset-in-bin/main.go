package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var data = gin.H{
	"Foo": "foo",
	"Bar": "bar",
}

func main() {
	r := gin.Default()
	t, err := loadTemplate()
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/foo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/template/foo.html", data)
	})
	r.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/template/bar.html", data)
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
