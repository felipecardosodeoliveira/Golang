package main

import (
	"fmt"
	"time"
)

func Task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	go Task("A")
	// Thread 3
	go Task("B")
	time.Sleep(15 * time.Second)
}
