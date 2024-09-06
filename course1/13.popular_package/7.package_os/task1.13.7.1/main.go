package main

import (
	"os"
	"path/filepath"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := WriteFile("./path/to/file.txt", []byte("Hello, World!\n"), os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
