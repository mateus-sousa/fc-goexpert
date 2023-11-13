package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string)

	// Thread 2
	go func() {
		// Estou preenchendo o canal com a string ola mundo
		canal <- "Ola mundo!"
	}()
	// Thread 1
	// Estou pegando o valor do canal e jogando para para variavel
	// O canal esvazia.
	msg := <-canal
	fmt.Println(msg)
}
