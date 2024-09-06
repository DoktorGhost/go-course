package main

import "fmt"

func main() {
	var num int = 10
	var str string = "Hello"

	fmt.Println(getVariableType(num))
	fmt.Println(getVariableType(str))
}

func getVariableType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}
