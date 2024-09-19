package main

type Director struct {
	builder PizzaBuilder
}

func newDirector(b PizzaBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b PizzaBuilder) {
	d.builder = b
}

func (d *Director) pizzaTime(size Size, thinDough, cheeseSide bool) Pizza {
	d.builder.SetSize(size)
	d.builder.SetThinDough(thinDough)
	d.builder.AddCheeseSide(cheeseSide)
	d.builder.SetSauce()
	d.builder.AddTopping()
	return d.builder.GetPizza()
}
