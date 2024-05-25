package main

import "fmt"

type ID int

var (
	a ID = 10
)

func main() {
	fmt.Println(a)
	fmt.Printf("-- value %v \n", a)
	fmt.Printf("-- type %T", a)
	fmt.Printf("-- type %d", a)

}
