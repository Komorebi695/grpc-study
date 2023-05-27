package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "grpcStudy/hello-server/proto"
	"net"
)

type server struct {
	pb.SayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 认证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("token 为空")
	}

	var appID string
	var appKey string

	if v, ok := md["appid"]; ok {
		appID = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	if appID != "nxj" || appKey != "123123" {
		return nil, errors.New("token 不正确")
	}

	return &pb.HelloResponse{
		ResponseMsg: "hello, " + req.RequestName,
	}, nil
}

func main() {
	// SSl认证
	// 参数：certFile, keyFile
	// 签名证书 和 私钥文件
	//creds, err := credentials.NewServerTLSFromFile("D:\\Programming\\ProjectCode\\GO\\src\\grpcStudy\\key\\test.pem",
	//	"D:\\Programming\\ProjectCode\\GO\\src\\grpcStudy\\key\\test.key")

	// 开启端口
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("fail: %v", err)
		return
	}
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	//grpcServer := grpc.NewServer(grpc.Creds(creds)) // ssl
	// 在grpc服务端中去注册自己写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}
}
