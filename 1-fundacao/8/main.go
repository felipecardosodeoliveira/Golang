package main

import "fmt"

func sum(arguments ...int) int {
	tot := 0
	for _, n := range arguments {
		tot += n
	}
	return tot
}

func main() {
	r := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("%d \n", r)
}
