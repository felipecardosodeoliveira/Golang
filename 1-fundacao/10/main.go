package main

import "fmt"

type Address struct {
	Street string
	Zip    string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c *Client) setActive(isActive bool) {
	c.Active = isActive
}

func main() {
	felipe := Client{
		Name:   "Felipe",
		Age:    31,
		Active: false,
	}

	fmt.Printf("Nome %s, Idade %d, Ativo %t \n", felipe.Name, felipe.Age, felipe.Active)

	felipe.setActive(true)

	fmt.Printf("Nome %s, Idade %d, Ativo %t \n", felipe.Name, felipe.Age, felipe.Active)

	felipe.Address.Street = "Terra"
	fmt.Println(felipe.Address.Street)
}
