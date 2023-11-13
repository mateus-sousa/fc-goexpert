package main

func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		// passa o valor true para o canal forever. liberando a thread principal que está ouvindo.
		forever <- true
	}()
	// Para a execução dessa thread para escutar o canal forever, e só seguira na execução da mesma quando o canal receber uma "carga"
	//Se o canal nunca receber carga o programa sofrerá um deadlock.
	<-forever
}
