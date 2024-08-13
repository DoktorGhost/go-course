package main

import "fmt"

func createCounter() func() int {
	count := 0
	return func() int {
		count = count + 1
		return count
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
