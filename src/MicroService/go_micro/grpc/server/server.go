package main

import (
	"fmt"
	"go_micro/grpc/myproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

// SayHello 函数
func (this *server) SayHello(ctx context.Context, in *myproto.HelloReq) (out *myproto.HelloRsp, err error) {
	return &myproto.HelloRsp{Msg: "hi," + in.Name + "!"}, nil
}

// SayName 函数
func (this *server) SayName(ctx context.Context, in *myproto.NameReq) (out *myproto.NameRsp, err error) {
	return &myproto.NameRsp{Msg: in.Name + ",luv u!"}, nil
}

func main() {

	// 创建网络
	listener, err := net.Listen("tcp", ":10010")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	// 创建gRPC服务
	srv := grpc.NewServer()

	// 注册服务
	myproto.RegisterHelloServerServer(srv, &server{})

	// 等待网络连接
	err = srv.Serve(listener)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}
