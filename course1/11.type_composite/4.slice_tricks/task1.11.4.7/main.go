package main

func Pop(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, []int{}
	}
	res := make([]int, len(xs)-1)
	resInt := xs[0]
	copy(res, xs[1:])
	return resInt, res
}
