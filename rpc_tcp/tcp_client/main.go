package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-14 20:26
// @description:

type Args struct {
	X, Y int
}

// 实现RPC 跨程序调用 -> 不在同一个内存空间

func main() {
	// 这里不一样了。
	client, err := rpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	// 异步调用，把返回结果通过通道方式获取到。看到那个go了吗，go
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
