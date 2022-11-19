package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc_exercise/public"
)

// @program:     rpc_exercise
// @file:        main.go
// @author:      bowen
// @create:      2022-11-14 21:54
// @description:

func main() {
	service := new(public.ServiceA)
	rpc.Register(service)

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, _ := l.Accept()
		//使用json协议
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
