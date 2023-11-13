package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64
	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{ID: i, Msg: "Message from RabbitMQ"}
			c1 <- msg
		}
	}()
	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{ID: i, Msg: "Message from Kafka"}
			c2 <- msg
		}
	}()
	for {
		select {
		case msg := <-c1: // RabbitMQ
			fmt.Printf("Mensagem recebida pelo RabbitMQ:%d, %s\n", msg.ID, msg.Msg)
		case msg := <-c2: // Kafka
			fmt.Printf("Mensagem recebida pelo Kafka:%d, %s\n", msg.ID, msg.Msg)
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout")
		}
	}
}
