package main

import "fmt"

func CountVowels(s string) int {
	vowelsMap := map[rune]struct{}{65: {}, 69: {}, 73: {}, 79: {}, 85: {}, 89: {}, 97: {}, 101: {}, 105: {}, 111: {}, 117: {}, 121: {}, 1072: {}, 1091: {}, 1086: {}, 1080: {},
		1101: {}, 1099: {}, 1103: {}, 1102: {}, 1077: {}, 1105: {}, 1040: {}, 1059: {}, 1054: {}, 1048: {}, 1069: {}, 1067: {}, 1071: {}, 1070: {}, 1045: {}, 1025: {}}
	count := 0
	runes := []rune(s)
	for _, run := range runes {
		if _, ok := vowelsMap[run]; ok {
			count++
		}
	}

	return count
}

func main() {
	count := CountVowels("Привет, Мир!")
	fmt.Println(count)

	count = CountVowels("Hello, World!")
	fmt.Println(count)
}
