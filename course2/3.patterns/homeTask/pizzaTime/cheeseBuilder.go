package main

type CheeseBuilder struct {
	pizza Pizza
}

func newCheeseBuilder() *CheeseBuilder {
	return &CheeseBuilder{}
}

func (b *CheeseBuilder) SetSize(size Size) {
	b.pizza.Size = size
}
func (b *CheeseBuilder) SetThinDough(dough bool) {
	b.pizza.ThinDough = dough
}

func (b *CheeseBuilder) AddCheeseSide(cheeseSide bool) {
	b.pizza.CheeseSide = cheeseSide
}

func (b *CheeseBuilder) SetSauce() {
	b.pizza.Sauce = AlfredoCheese
}

func (b *CheeseBuilder) AddTopping() {
	b.pizza.Toppings = []Topping{}
	b.pizza.Toppings = append(b.pizza.Toppings, Parmesan, Cheddar, Mozzarella)
}

func (b *CheeseBuilder) GetPizza() Pizza {
	return b.pizza
}
