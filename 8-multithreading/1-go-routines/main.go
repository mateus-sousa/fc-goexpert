package main

import (
	"fmt"
	"time"
)

// A main ja é nossa thread 1
func main() {
	go task("A")
	go task("B")
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Printf("A task anonima está rodando.\n")
			time.Sleep(1 * time.Second)
		}
	}()
	// Colocamos o sleep para dar tempo para que todas nossas threads paralelas fossem executadas antes que o programa fosse finalizado.
	time.Sleep(15 * time.Second)
}

func task(name string) {
	for i := 1; i < 10; i++ {
		fmt.Printf("A task %s está rodando.\n", name)
		time.Sleep(1 * time.Second)
	}
}
