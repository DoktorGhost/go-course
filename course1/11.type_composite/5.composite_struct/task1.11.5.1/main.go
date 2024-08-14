package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"strings"
	"time"
)

type User struct {
	Name string
	Age  int
}

func (u *User) sprintInfo() string {
	return fmt.Sprintf("Имя: %s, Возраст:=%d\n", u.Name, u.Age)
}

func getUsers() []User {
	gofakeit.Seed(time.Now().UnixNano())

	var result []User
	for i := 0; i < 10; i++ {
		user := User{Name: gofakeit.Name(), Age: gofakeit.Number(0, 100)}
		result = append(result, user)
	}
	return result
}

func preparePrint(users []User) string {
	var sb strings.Builder
	for _, user := range users {
		sb.WriteString(user.sprintInfo())
	}
	result := sb.String()
	return result
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}
