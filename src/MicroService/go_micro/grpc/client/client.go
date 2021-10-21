package main

import (
	"fmt"
	"go_micro/grpc/myproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 客户端连接服务器
	conn, err := grpc.Dial("127.0.0.1:10010", grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	// 关闭网络连接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("err:", err.Error())
		}
	}(conn)

	// 获取gRPC句柄
	client := myproto.NewHelloServerClient(conn)

	// 通过句柄调用函数
	rsp1, err := client.SayHello(context.Background(), &myproto.HelloReq{Name: "ohiua"})
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("SayHello:", rsp1.Msg)

	rsp2, err := client.SayName(context.Background(), &myproto.NameReq{Name: "shinrin"})
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println("SayName:", rsp2.Msg)
}
