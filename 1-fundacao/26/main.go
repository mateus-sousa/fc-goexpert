package main

import "fmt"

func main() {
	a := 2
	b := 1
	c := 3
	if a > b || c > a {
		fmt.Println("a > b e c> a")
	}
	switch a {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	case 3:
		fmt.Println("c")
	default:
		fmt.Println("d")
	}
}
