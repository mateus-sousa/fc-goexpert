package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()
	time.Sleep(time.Second * 3)

	select {
	case msg := <-c1:
		fmt.Println("recebido", msg)
	case msg := <-c2:
		fmt.Println("recebido", msg)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")
		//default:
		//	fmt.Println("Default")
	}
}
