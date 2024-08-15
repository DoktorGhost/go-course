package main

import (
	"fmt"
	"strings"
)

func createUniqueText(text string) string {
	result := make(map[string]struct{})
	var sb strings.Builder

	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		if _, ok := result[word]; !ok {
			result[word] = struct{}{}
			sb.WriteString(word)
			sb.WriteString(" ")
		}

	}
	resString := strings.TrimSpace(sb.String())
	return resString
}

func main() {
	fmt.Println(createUniqueText("bar bar bar foo foo baz"))
}
