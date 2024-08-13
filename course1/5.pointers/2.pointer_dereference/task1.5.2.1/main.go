package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := Dereference(&a)
	d := Sum(&b, &c)
	fmt.Println(c)
	fmt.Println(d)
}

func Dereference(n *int) int {
	res := *n
	return res
}

func Sum(a, b *int) int {
	res := *a + *b
	return res
}
