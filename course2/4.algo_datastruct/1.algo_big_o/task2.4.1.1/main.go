package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	if n < 1 {
		return 0
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster() bool {

	time1 := time.Now()
	factorialRecursive(100000)
	delta1 := time.Since(time1)

	time2 := time.Now()
	factorialIterative(100000)
	delta2 := time.Since(time2)

	if delta1 < delta2 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)
	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}
