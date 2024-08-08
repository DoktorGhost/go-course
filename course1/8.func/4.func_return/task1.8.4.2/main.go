package main

import "fmt"

func FindMaxAndMin(n ...int) (int, int) {
	maxInt := n[0]
	minInt := n[0]

	for _, i := range n {
		if i < minInt {
			minInt = i
		}
		if i > maxInt {
			maxInt = i
		}
	}
	return maxInt, minInt
}

func main() {
	maxInt, minInt := FindMaxAndMin(1, 3, 5, 1, 4, 5, 9, -4, 2, -10, 47, 75)
	fmt.Printf("maxInt: %d, minInt: %d\n", maxInt, minInt)
}
