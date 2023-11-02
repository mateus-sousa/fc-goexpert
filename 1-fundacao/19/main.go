package main

import "fmt"

// Type assertion, quando uma variavel é uma interface vazia, eu posso 'forçar' o seu uso como o tipo que eu desejo
// Porem não posso pedir para que um numero inteiro vire uma string, então nesse caso devo fazer uma verifação para validar se o type assertion deu certo
func main() {
	var myVar interface{} = "Hello World"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("o valor de res é %v e o resultado de ok é %v", res, ok)
}
