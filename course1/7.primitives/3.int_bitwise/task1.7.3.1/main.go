package main

import "fmt"

func main() {
	a := 60
	b := 13
	fmt.Printf(
		"a & b = %d\na | b = %d\na ^ b = %d\na << b = %d\na >> b = %d", bitwiseAnd(a, b),
		bitwiseOr(a, b), bitwiseXor(a, b), bitwiseLeftShift(a, b), bitwiseRightShift(a, b))
}

func bitwiseAnd(x, y int) int {
	return x & y
}
func bitwiseOr(x, y int) int {
	return x | y
}
func bitwiseXor(x, y int) int {
	return x ^ y
}
func bitwiseLeftShift(x, y int) int {
	return x << y
}
func bitwiseRightShift(x, y int) int {
	return x >> y
}
