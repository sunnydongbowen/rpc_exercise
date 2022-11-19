package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc_exercise/public"
)

// @program:     rpc_exercise
// @file:        main.go.go
// @author:      bowen
// @create:      2022-11-14 16:23
// @description: 七米视频，rpc server

func main() {

	service := new(public.ServiceA)
	rpc.Register(service) // 注册rpc服务

	rpc.HandleHTTP() // 基于http协议

	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	// 这里之前遇到过，先监听，后启动，大明老师有讲过
	http.Serve(l, nil)
}
