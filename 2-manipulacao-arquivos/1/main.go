package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Escrevendo bytes dentro de um arquivo usando go"))
	//size, err := f.WriteString("Escrevendo string")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", size)
	f.Close()

	// leitura direta
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// leitura de pouco em pouco
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
