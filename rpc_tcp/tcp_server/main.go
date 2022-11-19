package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_exercise/public"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-14 20:26
// @description: 基于tcp的rpc，看着和http很像，但其实不一样，这里可以回去看一下tcp时的代码！

func main() {
	service := new(public.ServiceA)
	rpc.Register(service)

	l, err := net.Listen("tcp", ":9091")

	if err != nil {
		log.Fatal("listen error:", err)
	}
	// 接收
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}

}
