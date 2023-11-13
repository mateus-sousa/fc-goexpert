package main

import (
	"fmt"
	"sync"
	"time"
)

// A main ja é nossa thread 1
func main() {
	//Vai ser o cara que irá impedir que nosso programa finalize antes que termine todas as tarefas das demais threads
	waitGroup := sync.WaitGroup{}
	// Define que ele deve esperar a execução de 25 tarefas.
	waitGroup.Add(25)
	go task("A", &waitGroup)
	go task("B", &waitGroup)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: A task anonima está rodando.\n", i)
			time.Sleep(1 * time.Second)
			// Define que uma tarefa foi executada, subtraindo um "credito" que foi definido no waitGroup.Add()
			waitGroup.Done()
		}
	}()
	// Mantera a execução do programa parada, liberando quando o contador de tarefas chegar a 0.
	waitGroup.Wait()
}

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: A task %s está rodando.\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}
