package main

import "fmt"

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	runes := []rune(*str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*str = string(runes)
}

func main() {
	a := 156
	mutate(&a)
	fmt.Println(a)

	str := "Съешь еще этих мягких французских булок, да выпей чаю"
	ReverseString(&str)
	fmt.Println(str)
}
