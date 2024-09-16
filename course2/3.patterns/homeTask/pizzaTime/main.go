package main

import "fmt"

func main() {
	builderCheese := newCheeseBuilder()
	builderChicken := newChickenBuilder()
	builderSea := newSeaBuilder()

	director := newDirector(builderCheese)
	pizza := director.pizzaTime(MediumSize, true, true)
	printPizzaInfo(pizza)

	director = newDirector(builderChicken)
	pizza = director.pizzaTime(BigSize, false, false)
	printPizzaInfo(pizza)

	director = newDirector(builderSea)
	pizza = director.pizzaTime(SmallSize, true, false)
	printPizzaInfo(pizza)

}

func printPizzaInfo(pizza Pizza) {
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Pizza info\n")
	fmt.Printf("-----------------------------\n")
	fmt.Printf("- Size: %s -\n", pizza.Size)
	if pizza.ThinDough {
		fmt.Printf("- Thin dough -\n")
	}
	fmt.Printf("- Sauce: %s -\n", pizza.Sauce)
	if pizza.CheeseSide {
		fmt.Printf("- Cheese Side -\n")
	}
	fmt.Printf("Toppings:\n")
	for i, v := range pizza.Toppings {
		fmt.Printf("-- %d) %v --\n", i+1, v)
	}
	fmt.Printf("-----------------------------\n\n")
}
