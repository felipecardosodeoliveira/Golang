package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel book canceled. Timeout reached.")
		return
	case <-time.After(time.Second * 2):
		fmt.Println("Hotel booked.")
	}
}
