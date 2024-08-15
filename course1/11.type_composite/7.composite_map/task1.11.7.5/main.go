package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	var sb strings.Builder

	words := strings.Fields(sentence)
	for _, word := range words {
		wordF := strings.ToLower(word)
		if filter[wordF] != true {
			sb.WriteString(word)
			sb.WriteString(" ")
		}

	}
	resString := strings.TrimSpace(sb.String())
	return resString
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}
