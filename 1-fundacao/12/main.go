package main

func main() {
	a := 10
	var p *int = &a
	println(a)
	*p = 11
	println(a)

	b := &a
	println(*b)

}
