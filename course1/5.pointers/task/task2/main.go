package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	p := &Person{Name: "Person 1"} // создали структуру Person с именем "Person 1" и передали указатель на эту структуру в переменную p
	modify(p)                      // в функцию передается указатель на структуру Person
	fmt.Println(p.Name)
}

func modify(person *Person) {
	person = &Person{Name: "Person 2"} //переменная person становится равна новому указателю на новую структуру Person со значением "Person 2"
	//но она определена только в стеке функциии modify
	// *person = Person{Name: "Person 2"}									//для модификации оригинала нужно разименовать ссылку на person и вней поменять значения Name
}
