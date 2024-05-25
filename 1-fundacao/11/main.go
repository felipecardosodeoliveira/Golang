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

type Company struct {
	Name string
}

type People interface {
	Desativar()
}

func (c *Company) Desativar() {

}

func (c *Client) Desativar() {
	c.Active = false
	fmt.Println("Desativado")
}

func Desativa(people People) {
	people.Desativar()
}

func main() {
	felipe := Client{
		Name:   "Felipe",
		Age:    31,
		Active: true,
	}

	fmt.Printf("Nome %s, Idade %d, Ativo %t \n", felipe.Name, felipe.Age, felipe.Active)

	Desativa(&felipe)

	fmt.Printf("Nome %s, Idade %d, Ativo %t \n", felipe.Name, felipe.Age, felipe.Active)

	empresa := Company{}

	fmt.Printf("Nome %s, \n", empresa.Name)

	Desativa(&empresa)

}
