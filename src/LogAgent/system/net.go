package system

import (
	"net"
	"strings"
)

func GetLocalIP() (ip string, err error) {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range address {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		ip = ipAddr.IP.String()
		return
	}
	return
}

func GetLocalIPByDial() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(addr.IP.String(), ":")[0]
	return
}
