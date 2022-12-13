package main

import (
	"github.com/hashicorp/consul/api"
	"log"
)

const (
	consulAddress = "192.168.72.130:8500" // 这个是连接的端口号
	port          = 8974                  // 这个是要注册的端口号
	serviceId     = "hello-72.130-8974"
	name          = "hello"
	address       = "192.168.72.130"
	//Tags =[]string{"NJ-hello", "hello", "bowen"}
)

func registerConsul() {
	// 连接consul
	config := api.DefaultConfig()
	config.Address = consulAddress
	// 这里本来传的是默认配置
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("api.Newclient failed,err:%v\n", err)
	}
	// 服务注册
	srv := &api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    name,
		Address: address,
		Port:    port,
		Tags:    []string{"NJ-hello", "hello", "bowen"},
	}
	client.Agent().ServiceRegister(srv)
}
