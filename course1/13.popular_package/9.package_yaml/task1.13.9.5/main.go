package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func main() {
	//data := []byte(`{"name":"Jhon", "age": 30}`)
	data := []byte(`
name: "Jhon"
age: 30
`)
	var person Person
	err := unmarshal(data, &person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Имя:", person.Name)
	fmt.Println("Возраст:", person.Age)
}

func unmarshal(data []byte, v interface{}) error {
	if json.Valid(data) {
		return json.Unmarshal(data, v)
	} else {
		return yaml.Unmarshal(data, v)
	}
}
