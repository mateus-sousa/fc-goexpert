package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := Client{
		Name:   "Mateus",
		Age:    27,
		Active: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t. \n", client.Name, client.Age, client.Active)
}
