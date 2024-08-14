package main

import "fmt"

func getSubSlice(xs []int, start, end int) []int {
	if start >= len(xs) || end >= len(xs) {
		return []int{}
	}
	if start >= end {
		return []int{}
	}
	return xs[start:end]
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subSlice := getSubSlice(numbers, 2, 6)
	fmt.Println(subSlice) // [3 4 5 6]
}
