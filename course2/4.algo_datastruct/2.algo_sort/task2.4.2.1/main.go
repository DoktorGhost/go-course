package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sort"
	"time"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %f, Count: %v, CreatedAt: %v\n", p.Name, p.Price, p.Count, p.CreatedAt)
}

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}

/*
//вариант 1 : создаем 3 структуры и в каждой из них реализуем по 3 метода len, less и swap


type ByPrice []Product

func (r ByPrice) Len() int {
	return len(r)
}

func (r ByPrice) Less(i, j int) bool {
	return r[i].Price < r[j].Price
}

func (r ByPrice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type ByCreatedAt []Product

func (r ByCreatedAt) Len() int {
	return len(r)
}

func (r ByCreatedAt) Less(i, j int) bool {
	return r[i].CreatedAt.Before(r[j].CreatedAt)
}

func (r ByCreatedAt) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type ByCount []Product

func (r ByCount) Len() int {
	return len(r)
}

func (r ByCount) Less(i, j int) bool {
	return r[i].Count < r[j].Count
}

func (r ByCount) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

*/

//вариант 2: создаем одну структуру, содержащую слайс Product и функцию сортировки, и для этой структуры пишем 3 метода.
//а дальше пишем 3 функции, чтобы для каждой сортировки использовалась своя функция

type ByFunc struct {
	products []Product
	lessFunc func(i, j int) bool
}

func (b ByFunc) Len() int {
	return len(b.products)
}
func (b ByFunc) Less(i, j int) bool {
	return b.lessFunc(i, j)
}
func (b ByFunc) Swap(i, j int) {
	b.products[i], b.products[j] = b.products[j], b.products[i]
}

func ByPrice(products []Product) ByFunc {
	return ByFunc{
		products: products,
		lessFunc: func(i, j int) bool {
			return products[i].Price < products[j].Price
		},
	}
}

func ByCreatedAt(products []Product) ByFunc {
	return ByFunc{
		products: products,
		lessFunc: func(i, j int) bool {
			return products[i].CreatedAt.Before(products[j].CreatedAt)
		},
	}
}

func ByCount(products []Product) ByFunc {
	return ByFunc{
		products: products,
		lessFunc: func(i, j int) bool {
			return products[i].Count < products[j].Count
		},
	}
}

func main() {
	products := generateProducts(10)
	fmt.Println("Исходный список:")
	fmt.Println(products)

	// Сортировка продуктов по цене
	sort.Sort(ByPrice(products))
	fmt.Println("\nОтсортировано по цене:")
	fmt.Println(products)

	// Сортировка продуктов по дате создания
	sort.Sort(ByCreatedAt(products))
	fmt.Println("\nОтсортировано по дате создания:")
	fmt.Println(products)

	// Сортировка продуктов по количеству
	sort.Sort(ByCount(products))
	fmt.Println("\nОтсортировано по количеству:")
	fmt.Println(products)
}
