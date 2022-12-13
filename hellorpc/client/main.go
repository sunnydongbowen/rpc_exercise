package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"rpc_exercise/hellorpc/client/pb"
	"strings"
	"time"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-22 16:34
// @description: 流式grpc以及consul发现服务等

func doRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 创建键值对
	md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	// 利用它生成一个context
	ctx = metadata.NewOutgoingContext(ctx, md)

	name := "bowen"
	// 把matedata传进去
	res, err := c.SayHello(ctx, &pb.Req{Name: name})
	if err != nil {
		log.Fatalf("s.Sayhello failed,err:%v\n", err)
	}
	log.Printf("got reply:%v\n", res.Reply)
}

func doServerStreamRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "bowen"
	//res, err := c.ServerStreamHello(ctx, &pb.Req{Name: name})

	//
	stream, err := c.ServerStreamHello(ctx, &pb.Req{Name: name})
	if err != nil {
		log.Fatalf("s.Sayhello failed,err:%v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("stream failed,err:%v\n", err)
		}
		// 把收到的每一次结果打出来
		log.Printf("got reply:%v\n", res.Reply)
	}
}

func doClientStreamRPC(c pb.HelloClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.ClientStreamHello(ctx)
	if err != nil {
		log.Fatalf("c.ClientStreamHello failed，err:%v\n", err)
	}

	names := []string{
		"hhh",
		"www",
		"aaa",
	}
	// 客户端流式发送请求
	for _, name := range names {
		stream.Send(&pb.Req{Name: name})
	}
	// 发送之后要告诉服务端并且接收响应
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv() failed,err:%v\n", err)
	}
	//打印响应结果
	log.Printf("got reply:%v\n", res.Reply)
}

func doBudiStreamRPC(c pb.HelloClient) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel() // 是这里退出的，不是第二个goroutine退出的。
	//err := timeoutCtx.Err()
	//fmt.Println("错误", err)

	stream, err := c.BudiStreamHello(timeoutCtx) // 这里的ctx要换一下
	if err != nil {
		log.Fatalf("BudiStreamHello failed,err: %v\n", err)
	}
	waitChan := make(chan struct{})
	// 开启单独的goroutine去接收消息
	go func() {
		defer func() {
			waitChan <- struct{}{}
		}()
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("Recv failed %v\n", err)
				return
			}
			//将收到是响应打印出来
			log.Printf("AI:%v\n", in.Reply)
		}
	}()

	go func() {
		// 一边收服务端响应，一边还有源源不断发送请求数据，要发送的数据是用户在终端输入的
		reader := bufio.NewReader(os.Stdin)
		for {
			c, _ := reader.ReadString('\n') // 没有输入就停在这了
			c = strings.TrimSpace(c)
			if len(c) == 0 {
				continue
			}

			if strings.ToUpper(c) == "EXIT" {
				break
			}
			stream.Send(&pb.Req{Name: c})
		}
		// 关闭发送流
		stream.CloseSend()
	}()

	<-waitChan
}

func main() {
	// 发现服务
	// 连接
	config := api.DefaultConfig()
	config.Address = "192.168.72.130:8500"
	consul, err := api.NewClient(config)
	if err != nil {
		fmt.Printf("NewClient　failed：%v\n", err)
	}
	//查询到可用
	m, err := consul.Agent().ServicesWithFilter("Service==`hello`") // 返回所有服务
	fmt.Println(m)
	if err != nil {
		fmt.Print(err)
	}
	var addr string
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)
		addr = fmt.Sprintf("%s:%d\n", v.Address, v.Port)
		fmt.Println(addr)
		// 使用第一个
		if len(addr) > 0 {
			break
		}
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial failed,err:%v\n", err)
	}
	defer conn.Close()
	// 创建rpc client
	client := pb.NewHelloClient(conn)
	log.Printf("client start")

	// 发送正常的rpc调用 一来一回
	doRPC(client)
	//doServerStreamRPC(client)
	//  doClientStreamRPC(client)
	//doBudiStreamRPC(client)

}
