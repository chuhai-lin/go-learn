package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	pb "proto/hello"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	// 连接
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)
}
