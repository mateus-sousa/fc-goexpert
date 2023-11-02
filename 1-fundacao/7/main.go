package main

import "fmt"

func main() {
	salaries := map[string]int{"Mateus": 1000, "Douglas": 2000, "Edilene": 3000, "Helena": 4000}

	for i, v := range salaries {
		fmt.Printf("O salario de %s é %d.\n", i, v)
	}

	// outras formas de declarar maps. PS: ja atribuindo um valor
	// sal1 := make(map[string]int)
	// sal2 := map[string]int{}
}
