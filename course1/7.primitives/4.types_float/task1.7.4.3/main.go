package main

import (
	"fmt"
	"math"
)

func CompareRoundedValues(a, b float64, decimalPlace int) (isEqual bool, difference float64) {
	round := func(value float64, places int) float64 {
		factor := math.Pow(10, float64(places))
		return math.Round(value*factor) / factor
	}

	aRounded := round(a, decimalPlace)
	bRounded := round(b, decimalPlace)

	isEqual = aRounded == bRounded

	difference = aRounded - bRounded

	return isEqual, difference

}

func main() {
	isEqual, difference := CompareRoundedValues(1.1234559, 1.1234771, 3)
	fmt.Println(isEqual, difference)
}
