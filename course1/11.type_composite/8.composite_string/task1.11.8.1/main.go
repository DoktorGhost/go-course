package main

func countBytes(s string) int {
	return len([]byte(s))
}

func countSymbol(s string) int {
	return len([]rune(s))
}
