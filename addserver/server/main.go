package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"rpc_exercise/addserver/server/proto"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-21 16:58
// @description:

// 定义一个结构体，满足能够注册到grpc的标准
// g
type server struct {
	proto.UnimplementedCalServer
}

// 主要的业务逻辑，其他部分其实都是固定格式
func (s *server) Do(ctx context.Context, in *proto.Req) (*proto.Res, error) {
	var res int64
	switch in.Op {
	case proto.Op_ADD:
		res = in.X + in.Y
	case proto.Op_SUB:
		res = in.X - in.Y
	default:
		res = 0
	}
	//sum := in.X + in.Y
	//return sum,这里让自己写，写的出来吗？
	return &proto.Res{Res: res}, nil
}
func main() {
	// 起tcp服务
	lis, err := net.Listen("tcp", ":8973")
	if err != nil {
		log.Fatalf("net.listen failed.err:%v\n", err)
	}

	// 注册rpc服务
	s := grpc.NewServer()
	proto.RegisterCalServer(s, &server{})

	//启动rpc服务
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("start server failed.err:%v\n", err)
	}

}
