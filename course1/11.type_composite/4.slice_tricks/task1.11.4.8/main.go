package main

func Shift(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, []int{}
	}
	first := xs[0]
	res := append(xs[(len(xs)-1):], xs[:(len(xs)-1)]...)
	return first, res
}
