package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Logger interface {
	Log(string) error
}

type RemoteLogger struct {
	Address string
}

func (rl *RemoteLogger) Log(msg string) error {

	client := &http.Client{}
	jsonStr := []byte(fmt.Sprintf(`{"message": "%s"}`, msg))

	req, err := http.NewRequest("POST", rl.Address, bytes.NewBuffer(jsonStr))
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	return nil
}

type FileLogger struct {
	File string
}

func (fl *FileLogger) Log(msg string) error {
	file, err := os.OpenFile(fl.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	if _, err := file.WriteString(currentTime + " : " + msg + "\n"); err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}

	return nil
}

type ConsoleLogger struct {
	Prefix string
}

func (cl *ConsoleLogger) Log(msg string) error {
	// Получаем текущее время и форматируем его
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Форматируем сообщение с датой, временем и префиксом (если он указан)
	if cl.Prefix != "" {
		msg = fmt.Sprintf("%s [%s]: %s", currentTime, cl.Prefix, msg)
	} else {
		msg = fmt.Sprintf("%s: %s", currentTime, msg)
	}

	// Выводим сообщение на консоль
	fmt.Println(msg)

	return nil
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}
func main() {
	consoleLogger := &ConsoleLogger{Prefix: "INFO"}
	remoteLogger := &RemoteLogger{Address: "https://google.com"}
	fileLogger := &FileLogger{File: "stdout.txt"} // Здесь замени на открытие реального файла, но для примера мы будем использовать os.Stdout
	loggers := []Logger{consoleLogger, fileLogger, remoteLogger}
	LogAll(loggers, "This is a test log message.")
}
