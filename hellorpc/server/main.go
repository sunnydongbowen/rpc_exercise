package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"rpc_exercise/hellorpc/server/pb"
	"strings"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-22 15:58
// @description: 多种流式grpc+consul

type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) SayHello(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	// 从客户端读取metadata.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if t, ok := md["timestamp"]; ok {
		fmt.Printf("timestamp from metadata:\n")
		// []string
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}
	return &pb.Res{
		Reply: "hello " + in.Name,
	}, nil
}

func (s *server) ServerStreamHello(in *pb.Req, stream pb.Hello_ServerStreamHelloServer) error {
	// 对传进来的人打招呼
	name := in.Name
	// 四国语言打招呼
	words := []string{"你好",
		"hello",
		"こんにちは",
		"여보세요"}
	for _, word := range words {
		if err := stream.Send(&pb.Res{Reply: word + name}); err != nil {
			log.Printf("stream send failed,err:%v\n", err)
			return err
		}
	}
	return nil
}

func (s *server) ClientStreamHello(stream pb.Hello_ClientStreamHelloServer) error {
	// 接收流式发来的请求数据
	var reply string = "你好哦"
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			//读完了后给客户端返回
			return stream.SendAndClose(&pb.Res{Reply: reply})
		}
		if err != nil {
			return err
		}
		reply += res.Name
	}
}

func (s *server) BudiStreamHello(stream pb.Hello_BudiStreamHelloServer) error {
	//服务端收一条 回复一条
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Printf("recv failed %v\n", err)
			return err
		}
		reply := magic(res.Name)
		//返回响应
		if err := stream.Send(&pb.Res{Reply: reply}); err != nil {
			log.Printf("stream.Send failed %v\n", err)
			return err
		}
	}
	return nil
}

// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

func main() {
	listener, err := net.Listen("tcp", ":8974")
	if err != nil {
		log.Fatalf("listened failed %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServer(s, &server{})
	//启动之前，注册服务到consul
	registerConsul()
	if err := s.Serve(listener); err != nil {
		log.Fatalf("served faile %v\n", err)
	}
}
