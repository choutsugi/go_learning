package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
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
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write failed, err:", err)
			continue
		}
	}
}
