package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpcStudy/hello-server/proto"
	"net"
)

type server struct {
	pb.SayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		ResponseMsg: "hello, " + req.RequestName,
	}, nil
}

func main() {
	// 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("fail: %v", err)
		return
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	// 在grpc服务端中去注册自己写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}
}
