package main

import (
	"fmt"
	"net/rpc"
)

var panda, result int

func main() {
	// 建立网络连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	panda = 1092

	// RPC远程调用
	err = client.Call("Panda.GetInfo", panda, &result)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	// 打印远程调用结果
	fmt.Println("result:", result)
}
