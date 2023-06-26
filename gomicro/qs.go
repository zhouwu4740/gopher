package gomicro

import (
	"context"
	"go-micro.dev/v4"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type HelloGomicro struct{}

func (h *HelloGomicro) Greeting(ctx context.Context, req *Request, resp *Response) error {
	resp.Message = "Hello " + req.Name
	return nil
}

func Quickstart() {
	service := micro.NewService(
		micro.Name("helloGomicro"),
		micro.Handle(new(HelloGomicro)),
		micro.Handle(":8088"),
	)
	service.Init()
	service.Run()
}
