package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGetReader(t *testing.T) {
	buffer := bytes.NewBufferString("Hello, World!")
	reader := getReader(buffer)

	if reader == nil {
		t.Error("getReader() вернула nil")
	}

	b := make([]byte, 13)
	n, err := reader.Read(b)
	if err != nil {
		t.Errorf("Ошибка при чтении из reader: %v", err)
	}

	if n != 13 {
		t.Errorf("Ожидалось 13 байтов, получено %d", n)
	}

	expected := "Hello, World!"
	if string(b) != expected {
		t.Errorf("Ожидалось '%s', получено '%s'", expected, string(b))
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)
	expected := "Hello, World!\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
