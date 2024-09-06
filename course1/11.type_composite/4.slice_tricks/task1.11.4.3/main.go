package main

func RemoveExtraMemory(xs []int) []int {
	result := make([]int, len(xs), len(xs))
	copy(result, xs)
	return result
}
