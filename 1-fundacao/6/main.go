package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	fmt.Printf("Tamanho: %d. Capacidade: %d. Itens: %v.\n", len(s), cap(s), s)

	// : antes do indice indica que estamos removendo os itens apos o indice informado
	// O : NÃO REMOVE OS ITENS DO SLICE ORIGINAL.
	fmt.Printf("Tamanho: %d. Capacidade: %d. Itens: %v.\n", len(s[:2]), cap(s[:2]), s[:2])

	// : depois do indice indica que estamos removendo os itens antes do indice informado
	fmt.Printf("Tamanho: %d. Capacidade: %d. Itens: %v.\n", len(s[4:]), cap(s[4:]), s[4:])

	s = append(s, 100)
	fmt.Printf("Tamanho: %d. Capacidade: %d. Itens: %v.\n", len(s), cap(s), s)
}
