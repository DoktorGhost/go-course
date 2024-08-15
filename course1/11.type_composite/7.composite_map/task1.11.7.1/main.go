package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWordOccurences(text string) map[string]int {
	result := make(map[string]int)

	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		word = strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		if word != "" {
			result[word] += 1
		}
	}

	return result
}

func main() {
	text := "Lorem ipsum, dolor sit amet! Consectetur adipiscing elit. Ipsum, lorem!"
	occurences := countWordOccurences(text)

	for word, count := range occurences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
