package main

import (
	"fmt"
	str "strings"
)

func ConcatenateStrings(sep string, strings ...string) string {
	var even []string
	var odd []string
	for i, v := range strings {
		if i%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return fmt.Sprintf("even: %s, odd: %s", str.Join(even, sep), str.Join(odd, sep))
}

func main() {
	fmt.Println(ConcatenateStrings("-", "hello", "world", "how", "are", "you"))
}
