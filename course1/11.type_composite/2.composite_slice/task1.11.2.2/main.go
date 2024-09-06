package main

import "fmt"

func MaxDifference(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return 0
	}
	maxx := numbers[0]
	minn := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > maxx {
			maxx = numbers[i]
		}
		if numbers[i] < minn {
			minn = numbers[i]
		}
	}
	return maxx - minn
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(MaxDifference(numbers))
}
