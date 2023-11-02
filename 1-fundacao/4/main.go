package main

import "fmt"

const a = "Hello, World"

type ID int

var (
	b bool    = true
	c string  = "Panda"
	d int     = 10
	e float64 = 1.5
	f ID      = 4
)

func main() {
	fmt.Printf("O tipo de f é %T e o valor é %v.", f, f)
}
