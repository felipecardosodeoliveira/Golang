package main

import "fmt"

func main() {
	var x interface{} = 10
	str, ok := x.(string)
	fmt.Printf("VALOR %v \n", x)
	fmt.Printf("Deu certo %v \n", ok)
	fmt.Printf("O tipo agora é %T \n", str)

}
