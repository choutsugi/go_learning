package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"micro/helloworld/handler"
	pb "micro/helloworld/proto"
)

const (
	// ServerName 服务名
	ServerName = "go.micro.srv.Helloworld"
)

func main() {
	// 创建服务
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Version("latest"),
	)

	// 注册句柄
	err := pb.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))
	if err != nil {
		logger.Fatal(err)
	}

	// 运行服务
	err = service.Run()
	if err != nil {
		logger.Fatal(err)
	}

}
