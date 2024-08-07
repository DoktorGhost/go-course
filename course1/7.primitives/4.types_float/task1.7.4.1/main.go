package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 6.52
	var b float64 = 5.43
	fmt.Println(hypotenuse(a, b))
}

func hypotenuse(a, b float64) float64 {
	square := math.Sqrt(a*a + b*b)
	return square
}
