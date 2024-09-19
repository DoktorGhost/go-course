package main

type Size string

const (
	SmallSize  Size = "Small"
	MediumSize Size = "Medium"
	BigSize    Size = "Big"
)

type Sauce string

const (
	SweetChili    Sauce = "Sweet Chili"
	AlfredoCheese Sauce = "Alfredo"
	CreamSauce    Sauce = "Cream sauce"
)

type Topping string

const (
	Mozzarella Topping = "Mozzarella"
	Cheddar    Topping = "Cheddar"
	Parmesan   Topping = "Parmesan"
	Chicken    Topping = "Chicken"
	Tomato     Topping = "Tomato"
	Shrimp     Topping = "Shrimp"
	Pineapple  Topping = "Pineapple"
)

type Pizza struct {
	Size       Size
	ThinDough  bool
	Sauce      Sauce
	CheeseSide bool
	Toppings   []Topping
}
