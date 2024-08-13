package main

import "fmt"

func main() {
	var name, city string
	var age int

	fmt.Println("Введите ваше имя:")
	fmt.Scanln(&name)
	fmt.Println("Введите ваш возраст:")
	fmt.Scanln(&age)
	fmt.Println("Введите ваш город:")
	fmt.Scanln(&city)

	fmt.Printf("Имя: %s\nВозрвст: %d\nГород: %s\n", name, age, city)

}
