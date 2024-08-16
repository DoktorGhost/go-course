package main

import (
	"fmt"
	"strings"
)

func concatStrings(xs ...string) string {
	var sb strings.Builder
	for _, x := range xs {
		sb.WriteString(x)
	}
	return sb.String()
}

func main() {
	result := concatStrings("Hello", " ", "World")
	fmt.Println(result)
}
