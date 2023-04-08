package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() // Colocar regra de expiração
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel() // Boa pratica sempre colcoar o cancel
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// ...
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
