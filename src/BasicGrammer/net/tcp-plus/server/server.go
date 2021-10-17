package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"zstone.com/tcp-plus/proto"
)

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close connection failed, err:", err)
		}
	}(conn)

	reader := bufio.NewReader(conn)

	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("Server recv:", msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("close listener failed, err:", err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go process(conn)
	}
}
