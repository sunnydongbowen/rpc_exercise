package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"rpc_exercise/grpc_server/pb"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-17 18:02
// @description: grpc server端

type server struct {
	pb.UnimplementedGreeterServer
}

func (S *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 业务逻辑
	var answer string
	if in.Name == "bowen" {
		answer = "加油"
	} else {
		answer = fmt.Sprintf("好好学习,%s是最棒的", in.Name)
	}
	return &pb.HelloReply{
		Answer: answer,
		Ts:     timestamppb.Now(),
	}, nil
}
func main() {
	// 向grpc注册我们的服务
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
		return
	}
	// 创建grpc服务
	s := grpc.NewServer()
	//  在gRPC服务端注册服务
	pb.RegisterGreeterServer(s, &server{})

	//在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(s)

	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve:%v", err)
		return
	}

}
