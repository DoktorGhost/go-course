package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func generateActivationKey() string {
	charset := "QWERTYUIOPASDFGHJKLZXCVBNM1234567890"
	char := []rune(charset)
	arr := []string{}

	for i := 0; i < 4; i++ {
		res := ""
		for j := 0; j < 4; j++ {
			res += string(char[rand.Intn(len(char))])
		}
		arr = append(arr, res)
	}

	return strings.Join(arr, "-")
}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey)
}
