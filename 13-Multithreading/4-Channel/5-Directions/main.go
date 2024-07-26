package main

import (
	"fmt"
)

// recebe apenas
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// envia apenas
func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go recebe("hello", hello)
	ler(hello)
}
