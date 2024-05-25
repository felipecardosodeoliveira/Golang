package main

import "fmt"

func main() {
	var t [3]string
	t[0] = "nome"
	t[1] = "sobrenome"
	t[2] = "endereco"

	// fmt.Println(t[0])
	// fmt.Println(t[len(t)-1])

	for i, v := range t {
		fmt.Printf("indice %d o valor %s \n", i, v)
	}
}
