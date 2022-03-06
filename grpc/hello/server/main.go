package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	pb "proto/hello"
)

type helloService struct{}

var HelloService = new(helloService)

const Address = "127.0.0.1:50052"

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	log.Println("Listen on " + Address)
	s.Serve(listen)
}
