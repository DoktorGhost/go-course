package main

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	if idx < 0 || idx >= len(xs) {
		return []int{}
	}
	half1 := make([]int, idx+1, idx+1)
	copy(half1, xs[:idx+1])
	half2 := make([]int, len(xs)-idx-1, len(xs)-idx-1)
	copy(half2, xs[idx+1:])

	half1 = append(half1, x...)
	half1 = append(half1, half2...)

	return half1
}
