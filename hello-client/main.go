package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpcStudy/hello-server/proto"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("D:\\Programming\\ProjectCode\\GO\\src\\grpcStudy\\key\\test.pem",
		"*.xiaojiang.com")
	if err != nil {
		return
	}

	//coon, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	coon, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds)) // ssl
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer coon.Close()

	// 建立连接
	client := pb.NewSayHelloClient(coon)

	// 执行rpc调用(这方法在服务端来实现并返回结果)
	rsp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		RequestName: "niexiaojiang",
	})
	if err != nil {
		return
	}

	fmt.Println(rsp.ResponseMsg)
}
