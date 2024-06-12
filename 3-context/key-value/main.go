package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "token", "senha")
	HotelBook(ctx)
}

func HotelBook(ctx context.Context) {
	fmt.Print(ctx)
}
