package main

import (
	"bufio"
	"fmt"
	"os"
)

// func ReadFile(p string) {
// 	b, err os.Open(p)
// 	if err {
// 		panic(err)
// 	}

// }

func main() {
	// f, err := os.Create("arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// size, err := f.Write([]byte("Olá mundo"))
	// size, err := f.WriteString("Olá mundo")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// fmt.Printf("tamanho do arquivo %d bytes \n ", size)

	// Lendo
	// f, err := os.ReadFile("arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	// lendo de forma parcial

	f, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	b := make([]byte, 2)
	for {
		n, err := r.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	os.Remove("arquivo.txt")
	// check(err)
}
