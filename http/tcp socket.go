package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	addr := "localhost:8081"
	tcpType := "tcp"
	TcpAddr, err := net.ResolveTCPAddr(tcpType, addr) //获取 tcpaddr
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	conn, err := net.DialTCP(tcpType, nil, TcpAddr)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	_, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	fmt.Println(string(result))
	os.Exit(0)
}
