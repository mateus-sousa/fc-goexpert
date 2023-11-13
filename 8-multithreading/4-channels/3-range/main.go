package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

// Vai rodar em thread 2
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		// a cada iteração publica o valor de i no canal ch.
		// caso o canal esteja cheio vai "travar" até que o canal seja esvaziado, para publicar o proximo valor.
		ch <- i
	}
	// Sinaliza que apos todas as publicações necessarias, o canal será fechado evitando um dealock pelo reader ficar esperando outra publicação eternamente.
	close(ch)
}

func reader(ch chan int) {
	// itera pelo canal, recebendo seus valores
	// assim que chega um novo valor no canal, ele esvazia o canal e joga seu valor na variavel x, e executa o que tem no corpo.
	for x := range ch {
		fmt.Printf("Recebido: %d\n", x)
	}
}
