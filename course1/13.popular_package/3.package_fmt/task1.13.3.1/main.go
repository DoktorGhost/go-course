package main

import (
	"fmt"
	"strconv"
	"strings"
)

func generateMathString(operands []int, operator string) string {
	if len(operands) < 2 {
		return "передан один операнд"
	}
	if operands[1] == 0 {
		return "делить на 0 нельзя"
	}
	if operator == "+" {
		res := 0
		arr := []string{}
		for _, i := range operands {
			res += i
			arr = append(arr, strconv.Itoa(i))
		}
		ss := strings.Join(arr, operator)
		return fmt.Sprintf("%s=%d", ss, res)
	} else if operator == "-" {
		res := operands[0]
		arr := []string{strconv.Itoa(res)}
		for i := 1; i < len(operands); i++ {
			res -= operands[i]
			arr = append(arr, strconv.Itoa(operands[i]))
		}
		ss := strings.Join(arr, operator)
		return fmt.Sprintf("%s=%d", ss, res)
	} else if operator == "*" {
		res := operands[0]
		arr := []string{strconv.Itoa(res)}
		for i := 1; i < len(operands); i++ {
			res *= operands[i]
			arr = append(arr, strconv.Itoa(operands[i]))
		}
		ss := strings.Join(arr, operator)
		return fmt.Sprintf("%s=%d", ss, res)
	} else if operator == "/" {
		res := float64(operands[0])
		arr := []string{strconv.Itoa(operands[0])}
		for i := 1; i < len(operands); i++ {
			res /= float64(operands[i])
			arr = append(arr, strconv.Itoa(operands[i]))
		}
		ss := strings.Join(arr, operator)
		return fmt.Sprintf("%s=%v", ss, res)
	}

	return "не передан аргумент"
}

func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "+"))
}
