package main

import "fmt"

func getFilePremissions(flag int) string {
	permission := map[int]string{
		0: "-,-,-",
		1: "-,-,Execute",
		2: "-,Write,-",
		3: "-,Write,Execute",
		4: "Read,-,-",
		5: "Read,-,Execute",
		6: "Read,Write,-",
		7: "Read,Write,Execute",
	}
	if flag == 0 {
		return fmt.Sprintf("Owner:%s Group:%s Other:%s\n", permission[0], permission[0], permission[0])
	}
	one := flag % 10
	two := (flag / 10) % 10
	three := (flag / 100) % 10

	return fmt.Sprintf("Owner:%s Group:%s Other:%s\n", permission[three], permission[two], permission[one])
}

func main() {
	fmt.Println(getFilePremissions(755))
	fmt.Println(getFilePremissions(0))
}
