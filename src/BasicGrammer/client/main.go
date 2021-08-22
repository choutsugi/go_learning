package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	reader := bufio.NewReader(os.Stdin)
	rBytes := make([]byte, 1024)

	for {
		wBytes, _, _ := reader.ReadLine()
		conn.Write(wBytes)
		n, _ := conn.Read(rBytes)
		fmt.Println("Receive from server: ", string(rBytes[0:n]))
	}
}
