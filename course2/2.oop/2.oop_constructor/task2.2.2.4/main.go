package main

import (
	"fmt"
	"io"
	"os"
)

// Logger interface
type Logger interface {
	Log(message string)
}

type LogSystem struct {
	logger Logger
}

func (l LogSystem) Log(message string) {
	l.logger.Log(message)
}

type LogOption func(*LogSystem)

func NewLogSystem(opts ...LogOption) *LogSystem {
	u := &LogSystem{}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

type FileLogger struct {
	file *os.File
}

func (fl FileLogger) Log(message string) {
	fl.file.WriteString(message)
	fl.file.Close()
}

type ConsoleLogger struct {
	out io.ReadWriter
}

func (cl ConsoleLogger) Log(message string) {
	fmt.Fprintln(cl.out, message)
}

func WithLogger(fl Logger) LogOption {
	return func(l *LogSystem) {
		l.logger = fl
	}
}

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()
	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))
	logSystem.Log("Hello, world!")
}
