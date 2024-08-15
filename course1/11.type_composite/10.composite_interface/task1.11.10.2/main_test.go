package main

import (
	"bytes"
	"os"
	"testing"
)

func TestOperateConcat(t *testing.T) {
	result := Operate(Concat, "Hello, ", "World!")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestOperateConcatMixed(t *testing.T) {
	result := Operate(Concat, "Hello, ", "World!", 159, 63)
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestOperateSumInts(t *testing.T) {
	result := Operate(Sum, 1, 2, 3, 4, 5)
	expected := 15
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestOperateSumFloats(t *testing.T) {
	result := Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)
	expected := 16.5
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestOperateEmpty(t *testing.T) {
	result := Operate(Concat)
	expected := ""
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestOperateWithNoMatchingType(t *testing.T) {
	result := Operate(Sum, "Hello", "World")
	expected := 0 // No numbers to sum, should be 0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
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
	expected := "Hello, World!\n15\n16.5\n"

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
