package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mateus-sousa/goexpert/5-packaging/3/math"
)

func main() {
	myMath := math.NewMath(1, 2)
	fmt.Println(myMath.Sum())
	fmt.Println(uuid.New().String())
}
