package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close connection failed, err:", err)
		}
	}(conn)

	reader := bufio.NewReader(conn)
	var buf [1024]byte

	for {
		rnum, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:rnum])
		fmt.Println("Server recv:", recvStr)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("close listener failed, err:", err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
