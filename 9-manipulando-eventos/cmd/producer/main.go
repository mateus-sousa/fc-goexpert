package main

import "github.com/mateus-sousa/fc-goexpert/9-manipulando-eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Publish(ch, "Ola palavra!", "amq.direct")
}
