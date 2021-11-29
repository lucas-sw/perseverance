package main

import (
	"fmt"
	"sync"
	"time"
)

type pass struct {
	RWM sync.RWMutex
	pwd string
}

var RoomPass = pass{pwd: "initPass"}

func Change(p *pass, pwd string) {
	defer p.RWM.Unlock()
	p.RWM.Lock()
	fmt.Println()
	time.Sleep(time.Second * 1)
	p.pwd = pwd
}

func getPwd(p *pass) string {
	defer p.RWM.RUnlock()
	p.RWM.RLock()
	fmt.Println("read pwd", p.pwd)
	time.Sleep(time.Second * 2)
	return p.pwd
}

func main() {

}
