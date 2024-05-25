package main

import "fmt"

func sum(a int, b int) (int, bool) {
	return a + b, false
}

func main() {
	r, f := sum(1, 2)
	fmt.Printf("%d %t\n", r, f)
}
