package main

import "fmt"

func main() {
	event := []string{"teste", "teste2", "teste3", "teste4"}
	event = append(event[:2], event[3:]...)
	fmt.Println(event)
}
