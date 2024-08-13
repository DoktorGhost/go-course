package main

import (
	"fmt"
	"strings"
)

func Factorial(n *int) int {
	nn := *n
	if nn < 0 {
		return 0
	}
	result := 1

	for i := 1; i <= nn; i++ {
		result *= i
	}
	return result
}

func isPalindrome(str *string) bool {
	s := *str
	s = strings.ToLower(s)

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	res := *numbers
	count := 0
	for _, i := range res {
		if i == *target {
			count++
		}
	}
	return count
}

func ReverseString(str *string) {
	runes := []rune(*str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*str = string(runes)
}

func main() {
	//Factorial
	a := 6
	fmt.Println(Factorial(&a))

	//isPalindrome
	s1 := "шалаш"
	s2 := "ШалАш"
	s3 := "не ПалинДром"

	fmt.Println(isPalindrome(&s1))
	fmt.Println(isPalindrome(&s2))
	fmt.Println(isPalindrome(&s3))

	//CountOccurrences
	arr := []int{1, 2, 5, 7, 9, 5, 1, 2, 3, 2, 1}
	target := 2
	fmt.Println(CountOccurrences(&arr, &target))

	//ReverseString
	fmt.Println(s1)
	ReverseString(&s1)
	fmt.Println(s1)

	fmt.Println(s2)
	ReverseString(&s2)
	fmt.Println(s2)

	fmt.Println(s3)
	ReverseString(&s3)
	fmt.Println(s3)
}
