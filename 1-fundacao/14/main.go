package main

import "fmt"

type Conta struct {
	saldo float32
}

func (c *Conta) soma(valor float32) float32 {
	c.saldo += valor
	return c.saldo
}

func main() {
	c := Conta{saldo: 100.}
	fmt.Printf("saldo %f \n", c.saldo)
	fmt.Printf("saldo somado %f \n", c.soma(10))
	fmt.Printf("saldo após soma %f \n", c.saldo)
}
