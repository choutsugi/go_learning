## TCP服务端

```go
func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.TCPConn){
	fmt.Println("client " + conn.RemoteAddr().String() + " has connected!")
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client " + conn.RemoteAddr().String() + " has disconnected!")
			break
		}
		fmt.Println(conn.RemoteAddr().String() + ": " + string(buf[0:n]))
		str := "Server get: " + string(buf[0:n])
		conn.Write([]byte(str))
	}
}
```

## TCP客户端

```go
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
```

