package main

import (
	"fmt"
	"sync"
	"time"
)

var bytesSlicePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	time1 := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		bytes := make([]byte, 1024)
		bytes[0] = 1
	}

	time2 := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		bytes := bytesSlicePool.Get().(*[]byte)
		(*bytes)[0] = 1
		bytesSlicePool.Put(bytes)
	}
	time3 := time.Now().Unix()
	fmt.Println("不使用缓存申请空间", time2-time1)
	fmt.Println("使用缓存申请空间", time3-time2)
}
