package main

import "fmt"

// Thread 1
func main() {
	cn := make(chan string) // vazio

	// Thread 2
	go func() {
		cn <- "Olá mundo" //cheio
	}()

	// Thread 1
	msg := <-cn // vazio
	fmt.Printf("%s \n", msg)
}
