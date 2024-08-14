package main

func Cut(xs []int, start, end int) []int {
	if start > len(xs)-1 || end > len(xs)-1 || start < 0 || end < 0 || start > end {
		return []int{}
	}
	res := make([]int, (end + 1 - start))
	copy(res, xs[start:end+1])
	return res
}
