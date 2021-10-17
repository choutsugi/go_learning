package main

import (
	"fmt"
	"net"

	"zstone.com/tcp-plus/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close connection failed, err:", err)
		}
	}(conn)

	for i := 0; i < 20; i++ {
		msg := "Hello Golang!"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("client write failed, err:", err)
			return
		}
	}
}
