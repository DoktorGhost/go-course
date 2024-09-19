package main

type ChickenBuilder struct {
	pizza Pizza
}

func newChickenBuilder() *ChickenBuilder {
	return &ChickenBuilder{}
}

func (b *ChickenBuilder) SetSize(size Size) {
	b.pizza.Size = size
}
func (b *ChickenBuilder) SetThinDough(dough bool) {
	b.pizza.ThinDough = dough
}

func (b *ChickenBuilder) AddCheeseSide(cheeseSide bool) {
	b.pizza.CheeseSide = cheeseSide
}

func (b *ChickenBuilder) SetSauce() {
	b.pizza.Sauce = CreamSauce
}

func (b *ChickenBuilder) AddTopping() {
	b.pizza.Toppings = []Topping{}
	b.pizza.Toppings = append(b.pizza.Toppings, Chicken, Mozzarella, Cheddar, Parmesan, Tomato)
}

func (b *ChickenBuilder) GetPizza() Pizza {
	return b.pizza
}
