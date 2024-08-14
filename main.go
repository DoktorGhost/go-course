package main

import (
	"fmt"
)

func main() {
	a := 1
	defer fmt.Println(&a) //1
	add(&a, 1)
	defer fmt.Println(&a) //2
	add(&a, 2)
	defer fmt.Println(&a) //4
	add(&a, 3)
	defer fmt.Println(&a) //7

}

func add(n *int, num int) {
	*n = *n + num
}
