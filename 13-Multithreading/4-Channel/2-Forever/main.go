package main

import "fmt"

// Thread 1
func main() {
	fv := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d \n", i)
			fv <- true
		}
	}()
	<-fv
}
