package main

func FilterDividers(xs []int, divider int) []int {
	if len(xs) == 0 || divider == 0 {
		return []int{}
	}

	var res []int

	for _, x := range xs {
		if x%divider == 0 {
			res = append(res, x)
		}
	}
	return res
}
