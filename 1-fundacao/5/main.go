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
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 15
	myArray[2] = 20

	for i, v := range myArray {
		fmt.Printf("O indice é %d e o valor é %d\n", i, v)
	}
}
