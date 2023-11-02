package main

import (
	"fmt"
	"github.com/mateus-sousa/goexpert/5-packaging/1/math"
)

func main() {
	myMath := math.NewMath(3, 4)
	fmt.Println(myMath.Sum())
}
