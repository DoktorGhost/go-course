package main

import "fmt"

func calculate(a, b int) (int, int, int, int, int) {
	sum := a + b
	difference := a - b
	product := a * b
	quotient := 0
	if b != 0 {
		quotient = a / b
	}
	remainder := a % b

	return sum, difference, product, quotient, remainder
}

func main() {
	var a, b int
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int
	a = 10
	b = 3
	sum, difference, product, quotient, remainder = calculate(a, b)
	fmt.Printf("a = %d b = %d sum = %d difference = %d product = %d quotient = %d remainder = %d", a, b, sum, difference, product, quotient, remainder)

}
