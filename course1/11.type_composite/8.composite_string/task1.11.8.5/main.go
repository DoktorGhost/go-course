package main

import "fmt"

func ReverseString(str string) string {
	strRune := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		strRune[i], strRune[j] = strRune[j], strRune[i]
	}
	return string(strRune)
}

func main() {
	fmt.Println(ReverseString("Hello, World"))
	fmt.Println(ReverseString("12345"))

}
