package main

import (
	"context"
	"fmt"
)

// Valor no context é muito pouco utilizado, um exemplo bom de uso é para passar dados de metricas e rastreabilidade da requisição
func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
