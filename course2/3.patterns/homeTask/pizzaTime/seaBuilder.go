package main

type SeaBuilder struct {
	pizza Pizza
}

func newSeaBuilder() *SeaBuilder {
	return &SeaBuilder{}
}

func (b *SeaBuilder) SetSize(size Size) {
	b.pizza.Size = size
}
func (b *SeaBuilder) SetThinDough(dough bool) {
	b.pizza.ThinDough = dough
}

func (b *SeaBuilder) AddCheeseSide(cheeseSide bool) {
	b.pizza.CheeseSide = cheeseSide
}

func (b *SeaBuilder) SetSauce() {
	b.pizza.Sauce = SweetChili
}

func (b *SeaBuilder) AddTopping() {
	b.pizza.Toppings = []Topping{}
	b.pizza.Toppings = append(b.pizza.Toppings, Shrimp, Pineapple, Mozzarella)
}

func (b *SeaBuilder) GetPizza() Pizza {
	return b.pizza
}
