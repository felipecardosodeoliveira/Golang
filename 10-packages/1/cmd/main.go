package main

import (
	"fmt"

	"teste.com/package/1/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Printf("A soma de %d + %d é %d \n", m.A, m.B, m.Add())
}
