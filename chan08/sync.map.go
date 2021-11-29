package main

import (
	"fmt"
	"sync"
)

var ma sync.Map

func main() {
	ma.Store(1, "one")
	ma.Store(1, "two")
	v, ok := ma.LoadOrStore(3, "three")
	fmt.Println(ok, v)
	v, ok = ma.LoadOrStore(1, "this one")
	fmt.Println(v, ok)
	v, ok = ma.Load(1)
	if ok {
		fmt.Println("key is existed, and value is:", v)
	} else {
		fmt.Println("key is not existed!")
	}
	f := func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	}
	ma.Range(f)
	ma.Delete(2)
	fmt.Println(ma.Load(2))
}
