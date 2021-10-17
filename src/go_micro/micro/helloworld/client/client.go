package main

import (
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"golang.org/x/net/context"
	helloworld "micro/helloworld/proto"
)

func main() {
	// 创建服务
	service := micro.NewService()

	service.Init()

	// 生成客户端
	cli := helloworld.NewHelloworldService("go.micro.srv.Helloworld", service.Client())

	// 调用远程方法
	rsp, err := cli.Call(context.Background(), &helloworld.Request{
		Name: "John",
	})

	if err != nil {
		logger.Info(err)
	}

	fmt.Println(rsp.Msg)
}
