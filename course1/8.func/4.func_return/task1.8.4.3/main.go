package main

import "fmt"

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	sum := price * float64(quantity)

	return sum, price
}

func main() {
	sum, price := CalculateStockValue(12.5, 150)
	fmt.Printf("Общая стоимость акций: %v, Цена одной акции: %v\n", sum, price)
}
