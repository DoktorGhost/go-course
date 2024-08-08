package main

import "fmt"

func Sum(a ...int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}

func Mul(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	result := 1
	for _, v := range a {
		result *= v
	}
	return result
}

func Sub(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	result := a[0]
	for _, v := range a[1:] {
		result -= v
	}
	return result
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Mul, 1, 7, 3))
	fmt.Println(MathOperate(Sub, 13, 2, 3))
}
