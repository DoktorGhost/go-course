package main

import "fmt"

func changeInt(a *int) {
	*a = 20
}

func changeFloat(b *float64) {
	*b = 6.28
}

func changeString(c *string) {
	*c = "Goodbye, world!"
}

func changeBool(d *bool) {
	*d = false
}

func main() {
	a := 12
	b := 7.52
	c := "Hello, World!"
	d := true

	fmt.Printf("a = %d, b = %v, c = %q, d = %v\n", a, b, c, d)

	changeInt(&a)
	changeFloat(&b)
	changeString(&c)
	changeBool(&d)

	fmt.Printf("a = %d, b = %v, c = %q, d = %v\n", a, b, c, d)
}
