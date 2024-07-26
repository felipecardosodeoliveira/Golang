package main

import (
	"fmt"
	"sync"
	"time"
)

func Task(name string, w *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		w.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(20)
	// Thread 2
	go Task("A", &waitGroup)
	// Thread 3
	go Task("B", &waitGroup)
	waitGroup.Wait()
}
