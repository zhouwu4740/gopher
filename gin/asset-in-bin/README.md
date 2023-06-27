# static assets
Go应用最终将编译成二进制文件，无法处理 html、css 等静态文件
利用 `go-assets`、`vfsgen` 第三方库可以解决此问题
参考 [go-assets]("https://gin-gonic.com/docs/examples/bind-single-binary-with-template/")
1. 利用`go-assets`解决 html 问题
```
github.com/jessevdk/go-assets
```
2. 利用`vfsgen`解决js/css/image等静态资源问题

## 安装
```shell
go get github.com/jessevdk/go-assets-builder
go install github.com/jessevdk/go-assets-builder
```

## 示例
1. 生成 assets.go，文件名称不限，重要的是生成的变量Assets名称
完成安装后，在应用根目录执行命令生成资产文件

```shell
go-assets-builder template -o assets.go
```


2. 实现Assets与 gin 集成
```go

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


```
