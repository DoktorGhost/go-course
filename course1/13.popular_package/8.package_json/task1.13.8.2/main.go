package main

import (
	"encoding/json"
	"fmt"
)

func getUsersFromJSON(data []byte) ([]User, error) {
	var users []User
	err := json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

type Comment struct {
	Text string `json:"text"`
}

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

func main() {
	jsonData := []byte(
		`[ 
			{
				"name": "Jhon",
				"age": 30,
				"comments": [
					{"text": "Great post!"},
					{"text": "I agree"}
				]
			},
			{
				"name": "Alice",
				"age": 25,
				"comments": [
					{"text": "Nice article"}
				]
			}
		]`)

	users, err := getUsersFromJSON(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Println("Name: ", user.Name)
		fmt.Println("Age: ", user.Age)
		fmt.Println("Comments: ")
		for _, comment := range user.Comments {
			fmt.Println("- ", comment.Text)
		}
		fmt.Println()
	}
}
