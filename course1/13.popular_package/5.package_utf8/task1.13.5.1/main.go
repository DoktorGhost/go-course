package main

import (
	"fmt"
	"unicode/utf8"
)

func countUniqueUTF8Chars(s string) int {
	uniqueChars := make(map[rune]bool)

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if r != utf8.RuneError {
			uniqueChars[r] = true
		}
		s = s[size:]
	}

	return len(uniqueChars)
}

func main() {
	str := "Hello,    !"
	fmt.Println(countUniqueUTF8Chars(str))
}
