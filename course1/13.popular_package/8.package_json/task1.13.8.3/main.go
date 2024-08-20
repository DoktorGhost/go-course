package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func writeJSON(filePath string, data []User) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
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
	users := []User{
		{
			Name: "Jhon",
			Age:  18,
			Comments: []Comment{
				{Text: "Grat Post"},
			},
		},
		{
			Name: "Jhon Uik",
			Age:  45,
			Comments: []Comment{
				{Text: "Its my dog"},
			},
		},
	}

	err := writeJSON("./users.json", users)
	if err != nil {
		fmt.Println(err)
		return
	}
}
