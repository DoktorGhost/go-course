package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	result := make(map[string]int, len(words))

	for _, word := range words {
		result[word] = 0
	}

	arr := strings.Fields(txt)

	for _, word := range arr {
		res := strings.ToLower(strings.Trim(word, ".,:!&=+-_()<>/"))
		if _, ok := result[res]; ok {
			result[res]++
		} else {
			continue
		}
	}
	return result
}

func main() {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris.         
Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.         
Praesent et diam eget libero egestas mattis sit amet vitae augue.
}`
	words := []string{"sit", "amet", "lorem"}
	result := CountWordsInText(txt, words)

	fmt.Println(result)

}
