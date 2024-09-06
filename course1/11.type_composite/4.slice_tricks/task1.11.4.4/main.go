package main

func RemoveIDX(xs []int, idx int) []int {
	if idx < 0 || idx >= len(xs) || len(xs) == 0 {
		return xs
	}

	return append(xs[:idx], xs[idx+1:]...)
}
