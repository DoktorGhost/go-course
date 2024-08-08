package main

import "fmt"

func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

func main() {
	fmt.Println(Factorial(6))
}
