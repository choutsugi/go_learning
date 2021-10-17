package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("connect udp server failed, err:", err)
		return
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("close socket failed, err:", err)
		}
	}(conn)

	sendData := []byte("Hello Golang!")
	_, err = conn.Write(sendData)
	if err != nil {
		fmt.Println("write msg to server failed, err:", err)
		return
	}

	data := make([]byte, 4096)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("recv data failed, err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
