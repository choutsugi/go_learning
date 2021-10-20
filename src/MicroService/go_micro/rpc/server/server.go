package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

/*
RPC包使用
*/

type Panda int

func (this *Panda) GetInfo(argType int, replyType *int) error {
	fmt.Println("Print:", argType)
	*replyType = argType + 1
	return nil
}

func pandaText(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "hello panda")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}

func main() {
	// panda请求
	http.HandleFunc("/panda", pandaText)

	// 类实例化为对象
	panda := new(Panda)
	// 服务端注册对象
	err := rpc.Register(panda)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	// 连接到网络
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":10010")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}
