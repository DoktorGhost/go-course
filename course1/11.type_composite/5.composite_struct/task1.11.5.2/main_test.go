package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGetAnimals(t *testing.T) {
	animals := getAnimals()

	if len(animals) != 10 {
		t.Errorf("Expected 10 animals, got %d", len(animals))
	}

	for _, animal := range animals {
		if animal.Type == "" || animal.Name == "" {
			t.Errorf("Expected non-empty animal type and name, got Type: '%s', Name: '%s'", animal.Type, animal.Name)
		}
		if animal.Age < 0 || animal.Age > 25 {
			t.Errorf("Expected age between 0 and 25, got %d", animal.Age)
		}
	}
}

func TestPreparePrint(t *testing.T) {
	animals := []Animal{
		{Type: "Dog", Name: "Buddy", Age: 5},
		{Type: "Cat", Name: "Whiskers", Age: 3},
	}

	result := preparePrint(animals)

	if len(result) == 0 {
		t.Errorf("Expected non-empty string, got empty string")
	}

	if !strings.Contains(result, "Тип: Dog, Имя: Buddy, Возраст:=5") {
		t.Errorf("Expected string to contain formatted output for Dog, got: %s", result)
	}

	if !strings.Contains(result, "Тип: Cat, Имя: Whiskers, Возраст:=3") {
		t.Errorf("Expected string to contain formatted output for Cat, got: %s", result)
	}
}

func TestSprintInfo(t *testing.T) {
	animal := Animal{Type: "Fish", Name: "Nemo", Age: 2}
	expected := "Тип: Fish, Имя: Nemo, Возраст:=2\n"

	result := animal.sprintInfo()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestMainFunction(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Вызываем main
	main()

	// Захватываем вывод
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Проверка
	output := buf.String()
	if !strings.Contains(output, "Тип:") {
		t.Errorf("Expected output to contain 'Тип:', got: %s", output)
	}
}
