package main

import "fmt"

func panic1() {
	panic("panic 2")
}

func panic2() {
	panic("panic 2")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic 1" {
				fmt.Println("Panic 1 recovered")
			}
		}
	}()
	fmt.Println("Hellow world")
	panic1()
}
