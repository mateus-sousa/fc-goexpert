package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(sum(&a, &b))
	fmt.Println(a)
	fmt.Println(b)
}

func sum(a, b *int) int {
	*a = 30
	*b = 40
	return *a + *b
}
