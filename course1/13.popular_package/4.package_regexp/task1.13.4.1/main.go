package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := "test@example.com"
	valid := isValidEmail(email)
	if valid {
		fmt.Printf("%s является валидным e-mail адресом\n", email)
	} else {
		fmt.Printf("%s не является валидным e-mail адресом\n", email)
	}
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)
	return re.MatchString(email)
}
