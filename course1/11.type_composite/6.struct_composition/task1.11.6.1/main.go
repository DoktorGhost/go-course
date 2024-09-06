package main

import "fmt"

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
	order.CalculateTotal()
}
func (order *Order) RemoveDish(dish Dish) {
	for i, dish2 := range order.Dishes {
		if dish.Name == dish2.Name {
			order.Dishes[i] = order.Dishes[len(order.Dishes)-1]
			order.Dishes = order.Dishes[:len(order.Dishes)-1]
			order.CalculateTotal()
			break
		}
	}
}
func (order *Order) CalculateTotal() {
	order.Total = 0
	for _, dish := range order.Dishes {
		order.Total += dish.Price
	}
}

func main() {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	order.RemoveDish(dish1)
	order.CalculateTotal()

	fmt.Println("Total:", order.Total)
}
