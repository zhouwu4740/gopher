package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func InjectLoggerEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf(`%s - [%s] "%s %s %s %d %s "%s" %s"\n`,
			p.ClientIP,
			p.TimeStamp.Format(time.DateTime),
			p.Method,
			p.Path,
			p.Request.Proto,
			p.StatusCode,
			p.Latency,
			p.Request.UserAgent(),
			p.ErrorMessage,
		)
	}))

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now()
		spanId := uuid.New()
		c.Set("spanId", spanId)
		c.Next()
		latency := time.Since(begin)
		log.Printf("[%s] latency = %d\n", spanId, latency)
		status := c.Writer.Status()
		log.Printf("[%s] status code = %d\n", spanId, status)
	}
}
