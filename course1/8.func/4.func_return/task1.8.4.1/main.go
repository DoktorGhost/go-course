package main

import "fmt"

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("check zero argument")
		}
	}()

	return a / b, a % b
}

func main() {
	divide, remainder := DivideAndRemainder(102, 0)
	fmt.Printf("Частное: %v, Остаток: %v.\n", divide, remainder)
}
