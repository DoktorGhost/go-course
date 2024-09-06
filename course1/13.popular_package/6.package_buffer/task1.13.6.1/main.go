package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func getReader(b *bytes.Buffer) *bufio.Reader {
	return bufio.NewReader(b)
}

func main() {
	buffer := bytes.NewBufferString("Hello, World!")
	b := make([]byte, 13)
	r := getReader(buffer)
	r.Read(b)
	fmt.Println(string(b))
}
