package main

import "fmt"

type Number interface {
	int | float32
}

func sum[T Number](m map[string]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	m := map[string]float32{"Kely": 10, "Felipe": 11.1}
	t := sum(m)
	fmt.Printf(" %v \n", t)
}
