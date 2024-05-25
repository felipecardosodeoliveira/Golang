package main

import "fmt"

func main() {
	salarios := map[string]int{"felipe": 1, "kely": 2, "Zisa": 3}
	for k, v := range salarios {
		fmt.Printf("o salario do(a) %s é %d \n", k, v)
	}
}
