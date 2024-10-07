package main

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

//go:generate easyjson -all $GOFILE
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

func EasyJSON(users []User) []User {
	var unmarshaledUsers []User
	for _, user := range users {
		data, err := easyjson.Marshal(user)
		if err != nil {
			continue
		}
		var userUnmarshaled User
		if err := easyjson.Unmarshal(data, &userUnmarshaled); err != nil {
			continue
		}
		unmarshaledUsers = append(unmarshaledUsers, userUnmarshaled)
	}
	return unmarshaledUsers
}

func JSON(users []User) []User {
	var unmarshaledUsers []User
	for _, user := range users {
		data, err := json.Marshal(user)
		if err != nil {
			continue
		}
		var userUnmarshaled User
		if err := json.Unmarshal(data, &userUnmarshaled); err != nil {
			continue
		}
		unmarshaledUsers = append(unmarshaledUsers, userUnmarshaled)
	}
	return unmarshaledUsers
}

func JSONiter(users []User) []User {
	var unmarshaledUsers []User
	for _, user := range users {
		data, err := jsoniter.Marshal(user)
		if err != nil {
			continue
		}
		var userUnmarshaled User
		if err := jsoniter.Unmarshal(data, &userUnmarshaled); err != nil {
			continue
		}
		unmarshaledUsers = append(unmarshaledUsers, userUnmarshaled)
	}
	return unmarshaledUsers
}

func GenerateUSER(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			ID:       i,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 14),
			Age:      gofakeit.Number(18, 100),
			Email:    gofakeit.Email(),
		}
	}
	return users
}
