package main

import "fmt"

// é um tipo derivado de um int
type MyNumber int

// uma constrint é escrita do conjunto de tipos que vai ser usado pelo generic
// o til permite que todos os tipos que "herdam" dos tipos usados pela constraint tbm sejam interpretados
type Number interface {
	~int | ~float64
}

// O tipo generic permite que criemos um tipo que "agrupa outros tipos"
func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m1 := map[string]int{
		"Jesus": 1000,
		"Maria": 2000,
		"Jose":  3000,
	}
	println(Sum(m1))
	m2 := map[string]float64{
		"Jesus": 1000.20,
		"Maria": 2000,
		"Jose":  3000,
	}
	fmt.Printf("%v", Sum(m2))
	m3 := map[string]MyNumber{
		"Jesus": 1000,
		"Maria": 2000,
		"Jose":  3000,
	}
	fmt.Printf("%v", Sum(m3))

	println(Compare("10", "10"))
}

// Questoes de comparação
func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}
