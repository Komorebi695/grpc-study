package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcStudy/hello-server/proto"
	"log"
)

type ClientTokenAuth struct{}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appID":  "nxj",
		"appKey": "123123",
	}, nil
}

// RequireTransportSecurity 是否需要基于TLS认证进行安全传输 false -> 不需要
func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//creds, err := credentials.NewClientTLSFromFile("D:\\Programming\\ProjectCode\\GO\\src\\grpcStudy\\key\\test.pem",
	//	"*.xiaojiang.com")
	//if err != nil {
	//	return
	//}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	//连接到server端，此处禁川安全传输，没有加密和验证
	coon, err := grpc.Dial("127.0.0.1:9090", opts...)
	//coon, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds)) // ssl
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
		fmt.Println(err.Error())
		return
	}

	fmt.Println(rsp.ResponseMsg)
}
