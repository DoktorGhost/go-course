package main

import "fmt"

// Пример кода на языке программирования
// Определение стека
var stack []int

// Функция для добавления элемента в стек
func push(value int) {
	stack = append(stack, value)
}

// Функция для удаления и возврата последнего элемента из стека
func pop() int {
	if len(stack) == 0 {
		fmt.Println("stack is empty")
		return 0
	}
	result := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return result
}
func main() {
	// Пример использования стека для операций
	push(5)
	push(3)
	result := pop() + pop()
	push(result)
	// Вывод результата выполнения программы
	fmt.Println(stack[0]) // 8
}
