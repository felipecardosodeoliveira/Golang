package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:1]), cap(s[:1]), s[:1])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	fmt.Printf("len=%d cap=%d %v\n", len(s[3:]), cap(s[3:]), s[3:])

	s = append(s, 60)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

}
