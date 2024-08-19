package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {
	buffer := bytes.NewBufferString("Hello, World!")

	result := getDataString(buffer)

	expected := "Hello, World!"
	if result != expected {
		panic(fmt.Sprintf("Expected %s, but got %s", expected, result))
	}
}

func getDataString(b *bytes.Buffer) string {
	reader := bufio.NewReader(b)
	result, _ := io.ReadAll(reader)
	return string(result)
}
