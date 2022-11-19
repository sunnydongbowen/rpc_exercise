package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-14 16:43
// @description: rpc client.客户端调用服务端的方法

// Args 定义一个Args参数类型
type Args struct {
	X, Y int
}

// 实现RPC 跨程序调用 -> 不在同一个内存空间

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dial failed:", err)
	}

	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("Service.Add: %d+%d=%d\n", args.X, args.Y, reply)
}
