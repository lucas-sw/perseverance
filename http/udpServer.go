package main

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello, I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func main() {
	p := make([]byte, 512)
	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 7778,
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some err", err)
		return
	}
	for {
		n, remoteAddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v : %s \n", remoteAddr, p[0:n])
		if err != nil {
			fmt.Printf("Error: %v", err)
			continue
		}
		go sendResponse(ser, remoteAddr)
	}
}
