package main

import "fmt"

func main() {
	hello := make(chan string)
	go receive("hello", hello)
	read(hello)
}

// Os <- no parametro do canal indica qual a direção ele pode utilizar nessa função
// o <- a direita do chan indica que ele é send only, podemos apenas enviar dados para o canal
func receive(nome string, hello chan<- string) {
	hello <- nome
}

// o <- a esquerda do chan indica que ele é receive only, podemos apenas receber dados do canal
func read(data <-chan string) {
	fmt.Println(<-data)
}
