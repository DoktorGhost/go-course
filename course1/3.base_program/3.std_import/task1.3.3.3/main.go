package main

import (
	"fmt"
	"math"
)

func main() {
	cos := Cos(1.57)
	sin := Sin(1.57)
	fmt.Printf("Cos: %f\nSin: %f\n", cos, sin)

}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}
