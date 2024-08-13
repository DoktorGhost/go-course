package main

import (
	"fmt"
	"math"
)

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		return price, fmt.Errorf("скидка не может превышать 50%%")
	}
	result := ((100 - discount) * price) / 100
	result = math.Round(result*100) / 100
	return result, nil
}

func main() {
	discountPrice, err := CheckDiscount(1452, 50.25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(discountPrice)
	}

	discountPrice, err = CheckDiscount(1452, 25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(discountPrice)
	}
}
