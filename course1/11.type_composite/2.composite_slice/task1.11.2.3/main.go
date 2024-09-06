package main

import "fmt"

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	uniqueNumber := 0
	for _, num := range numbers {
		uniqueNumber = bitwiseXOR(uniqueNumber, num)
	}
	return uniqueNumber
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) //5
}
