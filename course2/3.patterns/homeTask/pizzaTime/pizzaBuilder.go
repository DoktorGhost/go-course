package main

type PizzaBuilder interface {
	SetSize(size Size)
	SetThinDough(dough bool)
	AddCheeseSide(cheeseSide bool)
	SetSauce()
	AddTopping()
	GetPizza() Pizza
}
