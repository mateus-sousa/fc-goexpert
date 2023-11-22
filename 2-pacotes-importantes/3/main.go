package main

import (
	"fmt"
)

func main() {
	fmt.Println("PRIMEIRA LINHA")
	defer fmt.Println("SEGUNDA LINHA")
	fmt.Println("TERCEIRA LINHA")
}
