package main

import (
	"context"
	"fmt"
	"time"
)

func testWCancel(t int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("testWCancel.Done", ctx.Err())
			return
		case e := <-time.After(time.Duration(t) * time.Second):
			fmt.Println("testWCancel:", e)
		}
	}
}

func testDeadline(t int) {
	ctx := context.Background()
	dl := time.Now().Add(time.Duration(1*t) * time.Second)
	ctx, cancel := context.WithDeadline(ctx, dl)
	defer cancel()
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("testWDeadline.Done", ctx.Err())
			return
		case e := <-time.After(time.Duration(t) * time.Second):
			fmt.Println("testWDeadline:", e)
		}
	}

}

func main() {
	t := 1
	testWCancel(t)
	testDeadline(t)
}
