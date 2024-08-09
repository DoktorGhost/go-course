package main

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibonacciRecurs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibonacciRecurs(n-1) + FibonacciRecurs(n-2)
}
