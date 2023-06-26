package grpc

import (
	"context"
	"fmt"
	pb "github.com/zhouwu4740/tobegopher/grpc/user/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	resp := pb.Response{}
	resp.Message = "Hello " + req.Name
	return &resp, nil
}

func QuickStart() {
	fmt.Println("hello gRPC...")

	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Greeter{})

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
