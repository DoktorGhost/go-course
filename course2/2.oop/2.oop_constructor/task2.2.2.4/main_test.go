package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFileLogger(t *testing.T) {
	// Создаем временный файл для тестов
	file, err := os.CreateTemp("", "log_test_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // удаляем файл после теста

	fileLogger := FileLogger{file: file}
	fileLogger.Log("test message")

	// Проверяем содержимое файла
	data, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	expected := "test message"
	if string(data) != expected {
		t.Errorf("unexpected file content: got %q, want %q", string(data), expected)
	}
}

func TestConsoleLogger(t *testing.T) {
	var buf bytes.Buffer
	consoleLogger := ConsoleLogger{out: &buf}
	consoleLogger.Log("test message")

	expected := "test message\n"
	if got := buf.String(); got != expected {
		t.Errorf("unexpected buffer content: got %q, want %q", got, expected)
	}
}

func TestLogSystem(t *testing.T) {
	var buf bytes.Buffer
	consoleLogger := ConsoleLogger{out: &buf}
	logSystem := NewLogSystem(WithLogger(consoleLogger))
	logSystem.Log("test message")

	expected := "test message\n"
	if got := buf.String(); got != expected {
		t.Errorf("unexpected buffer content: got %q, want %q", got, expected)
	}
}

func TestNewLogSystemWithOptions(t *testing.T) {
	// Проверяем создание лог-системы с опциями
	file, err := os.CreateTemp("", "log_test_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name()) // удаляем файл после теста

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(&fileLogger))
	logSystem.Log("test message")

	// Проверяем содержимое файла
	data, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	expected := "test message"
	if string(data) != expected {
		t.Errorf("unexpected file content: got %q, want %q", string(data), expected)
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
	expected := ""

	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
