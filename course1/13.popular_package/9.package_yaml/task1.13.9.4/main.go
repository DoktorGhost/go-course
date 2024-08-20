package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func writeYAML(filePath string, data interface{}) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dataYAML, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(dataYAML)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Name     string   `yaml:"name"`
	Age      int      `yaml:"age"`
	Comments []Coment `yaml:"comments"`
}

type Coment struct {
	Text string `yaml:"text"`
}

func main() {
	users := []User{
		{
			Name: "Ivan23",
			Age:  42,
			Comments: []Coment{
				{
					Text: "Hello, my friends!",
				},
				{
					Text: "Yes, I do",
				},
			},
		},
		{
			Name: "Sergey",
			Age:  29,
			Comments: []Coment{
				{
					Text: "My name is Sergey",
				},
				{
					Text: "I from Russia",
				},
				{
					Text: "No, I Russian",
				},
			},
		},
	}

	err := writeYAML("./yaml/users.yaml", users)
	if err != nil {
		panic(err)
	}
}
