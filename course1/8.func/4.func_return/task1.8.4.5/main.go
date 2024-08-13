package main

import (
	"fmt"
	"math"
)

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	result := ((finalValue - initialValue) * 100) / initialValue
	result = math.Round(result*100) / 100
	return result
}

func main() {
	percent := CalculatePercentageChange(2233.36578, 3242.432)
	fmt.Printf("%v%%", percent)
}
