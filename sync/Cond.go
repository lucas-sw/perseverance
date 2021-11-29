package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ready     = false
	singerNum = 3
)

func Sing(singerId int, c *sync.Cond) {
	fmt.Printf("Singer (%d) is ready\n", singerId)
	c.L.Lock()
	for !ready {
		fmt.Printf("Singer (%d) is waiting\n", singerId)
		c.Wait()
	}
	fmt.Printf("Singer(%d)is sing a song\n", singerId)
	//ready = false
	c.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < singerNum; i++ {
		go Sing(i, cond)
	}
	time.Sleep(time.Second * 3)
	for i := 0; i < singerNum; i++ {
		ready = true
		//cond.Broadcast()
		cond.Signal()
		time.Sleep(time.Second * 3)
	}
}
