package main

import "fmt"

type Address struct {
	District   string
	Street     string
	PostalCode string
	Number     string
}

type Person interface {
	Inactivate()
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c *Client) Inactivate() {
	c.Active = false
}

func main() {
	client := Client{
		Name:   "Mateus",
		Age:    27,
		Active: true,
	}
	Inativation(&client)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t. \n", client.Name, client.Age, client.Active)
}

func Inativation(person Person) {
	person.Inactivate()
}
