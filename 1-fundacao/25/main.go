package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	positions := []string{"primeiro", "segundo", "terceiro"}
	for k, v := range positions {
		fmt.Println(k, v)
	}
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
