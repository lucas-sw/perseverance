package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error :", err.Error())
		os.Exit(1)
	}
}

func nowtime() string {
	return time.Now().String()
}

func main() {
	addr := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkError(err)
	mylistener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	i := 0
	for {
		myConn, err := mylistener.Accept()
		fmt.Printf("myconn ")
		if err != nil {
			continue
		}
		i++
		nowTime := nowtime()
		fmt.Printf("request no %d return time:%s \n", i, nowTime)
		myConn.Write([]byte(nowTime))
		myConn.Close()
	}
}

// 打开 terminal , telnet localhost 7777 进行测试
