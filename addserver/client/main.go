package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"rpc_exercise/addserver/client/proto"
	"time"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-21 20:34
// @description:

func main() {
	// 通过rpc调用其他程序的Do方法

	// 1. 建立链接
	conn, err := grpc.Dial(":8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial failed:%v\n", err)
	}
	defer conn.Close()

	// 2.发起调研
	client := proto.NewCalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//res, err := client.Do(ctx, &proto.Req{X: 10, Y: 20, Op: proto.Op_SUB})
	res, err := client.Do(ctx, &proto.Req{X: 10, Y: 20})
	if err != nil {
		log.Fatalf("call failed:%v\n", err)
	}
	fmt.Println(res.Res)
}
