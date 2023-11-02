package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sum(51, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func sum(a, b int) (int, error) {
	if (a + b) > 50 {
		return 0, errors.New("Resultado é maior que 50")
	}
	return a + b, nil
}
