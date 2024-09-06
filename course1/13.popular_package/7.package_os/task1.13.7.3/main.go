package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadString(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Sprintf("Error opening file: %v", err)
	}
	defer file.Close()
	result := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"

	}
	return result
}

func main() {
	filePath := "./course1/13.popular_package/7.package_os/task1.13.7.2/file.txt"
	str := ReadString(filePath)
	fmt.Println(str)
}
