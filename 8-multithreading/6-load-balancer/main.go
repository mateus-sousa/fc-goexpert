package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	workersAmount := 1000
	for i := 0; i < workersAmount; i++ {
		go worker(i, data)
	}
	for i := 0; i < 10000; i++ {
		data <- i
	}
}

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
