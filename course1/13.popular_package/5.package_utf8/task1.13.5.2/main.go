package main

import (
	"fmt"
	"strings"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	ss := strings.ToLower(s)
	for _, char := range ss {
		if isRussianLetter(char) {
			counts[char]++
		}
	}
	return counts
}

func isRussianLetter(char rune) bool {
	if char >= 1040 && char <= 1103 {
		return true
	}
	return false
}

func main() {
	result := countRussianLetters("Привет, Мир!")
	for key, value := range result {
		fmt.Printf("%c: %d ", key, value)
	}
}
