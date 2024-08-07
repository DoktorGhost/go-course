package main

import (
	"fmt"
	"math"
)

var CalculateCircleArea func(radius float64) float64
var CalculateRectangleArea func(width, height float64) float64
var CalculateTriangleArea func(base, height float64) float64

func main() {
	CalculateCircleArea = func(radius float64) float64 {
		return radius * radius * math.Pi
	}
	CalculateRectangleArea = func(width, height float64) float64 {
		return width * height
	}
	CalculateTriangleArea = func(base, height float64) float64 {
		return base * height * 0.5
	}

	fmt.Printf("Circle: %.2f\n", CalculateCircleArea(5.0))
	fmt.Printf("Rectangle: %.2f\n", CalculateRectangleArea(4.0, 6.0))
	fmt.Printf("Triangle: %.2f\n", CalculateTriangleArea(3.0, 7.0))
}
