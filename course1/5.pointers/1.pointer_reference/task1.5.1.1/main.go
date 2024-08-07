package main

import "fmt"

func Add(a, b int) *int {
	c := a + b
	return &c
}

func Max(numbers []int) *int {
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return &max
}

func IsPrime(number int) *bool {
	var result bool
	if number <= 1 {
		result = false
	}
	if number == 2 {
		result = true
	}
	if number%2 == 0 {
		result = false
	} else {
		result = true
		for i := 3; i*i <= number; i += 2 {
			if number%i == 0 {
				result = false
				break
			}
		}
	}

	return &result
}

func ConcatenateStrings(strs []string) *string {
	var result string
	for _, str := range strs {
		result += str
	}
	return &result
}

func main() {
	//add
	a := 100
	b := 150
	res := Add(a, b)
	fmt.Println(res)
	fmt.Println(*res)

	//max
	arr := []int{13, 4, 13, 12, 98, 76}
	res = Max(arr)
	fmt.Println(res)
	fmt.Println(*res)

	//isPrime
	c := 13
	d := 71
	e := 64

	flag := IsPrime(c)
	fmt.Println(flag)
	fmt.Println(*flag)

	flag = IsPrime(d)
	fmt.Println(flag)
	fmt.Println(*flag)

	flag = IsPrime(e)
	fmt.Println(flag)
	fmt.Println(*flag)

	//ConcatenateStrings

	arrStr := []string{"Hello", ",", " ", "World", "!"}
	resStr := ConcatenateStrings(arrStr)
	fmt.Println(resStr)
	fmt.Println(*resStr)

}
