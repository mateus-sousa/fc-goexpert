package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// cria um contexto vazio em background
	ctx := context.Background()
	// adiciona uma "condição de cancelamento para o context, de timeout
	// passado 3 segundos o contexto será cancelado
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	// É uma boa pratica SEMPRE cancelar o contexto ao final do processamento da aplicação
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// O select executara o case do channel que receber uma sinalização primeiro
	select {
	// No momento que o context for cancelado, o channel ctx.Done receberá uma sinalização
	// E printará o que está no case
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	// Após o tempo passado por parametro, o channel time.After receberá uma sinalização
	// E printará o que está no case, neste caso este case sera printado pois o tempo de espera é menor que o timeout do cancel do ctx.
	case <-time.After(time.Second * 2):
		fmt.Println("Hotel booked")
	}
}
