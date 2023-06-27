package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var graceTime = time.Second * 10

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		time.Sleep(time.Second * 10)
		c.String(http.StatusOK, "Welcome, gopher!")
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cfunc := context.WithTimeout(context.Background(), graceTime)
	defer cfunc()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown Server: ", err)
	}

	select {
	case <-ctx.Done():
	}

	log.Println("Server exit")
}
