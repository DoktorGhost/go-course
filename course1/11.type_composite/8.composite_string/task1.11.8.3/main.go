package main

import "fmt"

func getBytes(s string) []byte {
	return []byte(s)
}

func getRunes(s string) []rune {
	return []rune(s)
}

func main() {
	s := "Hello, МИР!"
	runes := getRunes(s)
	fmt.Println(runes)
	b := getBytes(s)
	fmt.Println(b)
}
