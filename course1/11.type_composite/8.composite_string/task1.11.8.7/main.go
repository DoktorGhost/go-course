package main

import "fmt"

func ReplaceSymbols(str string, old, new rune) string {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if runes[i] == old {
			runes[i] = new
		}
	}
	return string(runes)
}

func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result)
}
