package main

import (
	"fmt"
	"strings"
)

func main() {
	// Test case 1
	result := UserInfo("John", 21, "Moscow", "Saint Petersburg")
	fmt.Println(result)

	// Test case 2
	result = UserInfo("Alex", 34)
	fmt.Println(result)
}

func UserInfo(name string, age int, cities ...string) string {
	return fmt.Sprintf("Имя: %s, возраст: %d, города: %v\n", name, age, strings.Join(cities, ", "))
}
