package main

import (
	"fmt"
)

func main() {
	result := func() int {
		return sum(51, 123, 123, 5, 12, 575, 88, 100) * 2
	}()
	fmt.Println(result)
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
