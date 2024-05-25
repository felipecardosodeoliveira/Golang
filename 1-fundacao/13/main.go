package main

func sum(a, b *int) int {
	*a = 25
	return *a + *b
}

func main() {
	a := 10
	b := 11
	println(sum(&a, &b))
	println(a)
}
