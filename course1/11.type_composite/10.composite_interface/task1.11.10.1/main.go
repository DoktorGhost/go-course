package main

import "fmt"

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case []int:
		return "[]int"
	case []float64:
		return "[]float64"
	case []string:
		return "[]string"
	case nil:
		return "Пустой интерфейс"
	default:
		return "Тип неопределен"
	}
}

func main() {
	var i interface{} = 42
	fmt.Println(getType(i))

	var j interface{} = "hello world"
	fmt.Println(getType(j))

	var k interface{} = []int{1, 2, 3}
	fmt.Println(getType(k))

	var l interface{} = interface{}(nil)
	fmt.Println(getType(l))
}
