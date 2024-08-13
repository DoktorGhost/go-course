package main

import (
	"fmt"
	"math"
)

func main() {
	result := Floor(12.3423)
	fmt.Println(result)

}

func Floor(x float64) float64 {
	return math.Floor(x)
}
