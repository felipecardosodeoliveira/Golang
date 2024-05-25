package main

import "fmt"

func main() {
	var x interface{} = 10

	showType(x)
}

func showType(t interface{}) {
	fmt.Printf("O tipo é %T e o valor é %v \n", t, t)
}
