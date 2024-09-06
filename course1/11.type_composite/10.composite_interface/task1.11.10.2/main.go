package main

import "fmt"

var Operate = func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

var Concat = func(xs ...interface{}) interface{} {
	var res string
	for _, x := range xs {
		if str, ok := x.(string); ok {
			res += str
		} else {
			continue
		}
	}
	return res
}

var Sum = func(xs ...interface{}) interface{} {
	var intSum int
	var floatSum float64

	for _, x := range xs {
		switch v := x.(type) {
		case int:
			intSum += v
		case float64:
			floatSum += v
		}
	}

	if floatSum > 0 {
		return floatSum
	}
	return intSum
}

func main() {
	fmt.Println(Operate(Concat, "Hello, ", "World!"))  //Hello, World!
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           //15
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) //16.5
}
