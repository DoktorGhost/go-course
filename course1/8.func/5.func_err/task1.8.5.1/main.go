package main

import (
	"fmt"
	"math"
	"strconv"
)

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	initialValueFloat, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, fmt.Errorf("некорректное число в запросе: %v", err)
	}

	finalValueFloat, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, fmt.Errorf("некорректное число в запросе: %v", err)
	}
	if initialValueFloat == 0 {
		return 0, nil
	}
	result := ((finalValueFloat - initialValueFloat) * 100) / initialValueFloat
	result = math.Round(result*100) / 100
	return result, nil
}

func main() {
	percent, err := CalculatePercentageChange("0", "10")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v%%", percent)
	}

}
