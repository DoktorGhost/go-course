package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	myPanic()
}

func myPanic() {
	panic("my panic message")
}

/*
//исходные данные
func main() {
	ch := make(chan string)
	//go myPanic(ch)  //но тут горутина
	myPanic(ch)
	myPanic()
}

func myPanic(ch chan string) {
	defer func() {
		if err := recover(); err != nil {
			ch <- fmt.Sprint(err)
		}
	}()
	panic("my panic message")
}
*/
