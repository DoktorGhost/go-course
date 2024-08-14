package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"strings"
	"time"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func (a *Animal) sprintInfo() string {
	return fmt.Sprintf("Тип: %s, Имя: %s, Возраст:=%d\n", a.Type, a.Name, a.Age)
}

func getAnimals() []Animal {
	gofakeit.Seed(time.Now().UnixNano())

	var result []Animal
	for i := 0; i < 10; i++ {
		user := Animal{Type: gofakeit.AnimalType(), Name: gofakeit.PetName(), Age: gofakeit.Number(0, 25)}
		result = append(result, user)
	}
	return result
}

func preparePrint(animals []Animal) string {
	var sb strings.Builder
	for _, animal := range animals {
		sb.WriteString(animal.sprintInfo())
	}
	result := sb.String()
	return result
}

func main() {
	animals := getAnimals()
	result := preparePrint(animals)
	fmt.Println(result)
}
