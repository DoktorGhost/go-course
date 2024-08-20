package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func writeJSON(filePath string, data interface{}) error {
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

func main() {
	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraster",
			"age":  30,
		},
	}

	err := writeJSON("./data.json", data)
	if err != nil {
		panic(err)
	}
}
