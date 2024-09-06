package main

func InsertToStart(xs []int, x ...int) []int {
	var result []int
	for _, xx := range x {
		result = append(result, xx)
	}
	result = append(result, xs...)
	return result
}
