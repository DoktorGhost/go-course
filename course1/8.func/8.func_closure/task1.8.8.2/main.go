package main

import "fmt"

func multiplier(factor float64) func(float64) float64 {
	return func(x float64) float64 {
		return x * factor
	}
}

func main() {
	m := multiplier(2.5)
	result := m(10)
	fmt.Println(result) //25
	result = m(100)
	fmt.Println(result) //250
}
