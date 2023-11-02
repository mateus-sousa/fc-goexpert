package main

import "fmt"

func main() {
	var first interface{} = 10
	var second interface{} = "Hello World"

	showType(first)
	showType(second)
}

func showType(t interface{}) {
	fmt.Printf("o tipo da variavel é %T e o valor é %v\n", t, t)
}
