package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	defer w.Done()
	for job := range jobs {
		fmt.Printf("worker(%d) start to do job (%d)\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker(%d) finished job(%d)\n", id, job)
		results <- job * job
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for id := 0; id < 3; id++ {
		w.Add(1)
		go worker(id, jobs, results)
	}
	go func() {
		for {
			fmt.Println(<-results)
		}
	}()
	for job := 0; job < 200; job++ {
		jobs <- job
	}
	close(jobs)
	w.Wait()

}
