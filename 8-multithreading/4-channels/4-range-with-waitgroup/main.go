package main

import (
	"fmt"
	"sync"
)

// Utilizando waitgroups podemos ter tanto o publisher quanto o reader em threads paralelas.
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

// Vai rodar em thread 2
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Recebido: %d\n", x)
		wg.Done()
	}
}
