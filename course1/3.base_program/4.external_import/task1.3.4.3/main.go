package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func main() {
	fmt.Println(GenerateFakeData())
}

func GenerateFakeData() string {
	name := fake.FirstName()
	surname := fake.LastName()
	address := fake.StreetAddress()
	phone := fake.Phone()
	email := fake.EmailAddress()

	return fmt.Sprintf("Name: %s %s\nAddres: %s\nPhone: %s\nEmail: %s\n", name, surname, address, phone, email)
}
