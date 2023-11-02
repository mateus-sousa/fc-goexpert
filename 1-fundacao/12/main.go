package main

import "fmt"

type Address struct {
	District   string
	Street     string
	PostalCode string
	Number     string
}

// Address compoe cliente, e podemos usar os atributos de address como se fossem de Client
type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	client := Client{
		Name:   "Mateus",
		Age:    27,
		Active: true,
	}

	client.District = "distrito"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t. \n", client.Name, client.Age, client.Active)
}
