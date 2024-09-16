package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// уровни логирования
const (
	DEBUG = iota // 0
	INFO         // 1
	ERROR        // 3
)

// цвет уровней логгирования, чисто для красоты
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
)

type Logger struct {
	logger *log.Logger
	level  int
}

var once sync.Once

var singleInstance *Logger

func GetInstance() *Logger {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &Logger{
					logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
					level:  INFO,
				}
				message := fmt.Sprintf("Успешное создание логгера")
				singleInstance.logMessage(ColorYellow, "INFO", message)
			})
	} else {
		message := fmt.Sprintf("Логгер был создан ранее")
		singleInstance.logMessage(ColorYellow, "INFO", message)
	}

	return singleInstance
}

func (l *Logger) SetLevel(level int) {
	l.level = level
	levelName := "INFO"
	if level == 0 {
		levelName = "DEBUG"
	} else if level == 2 {
		levelName = "ERROR"
	}
	message := fmt.Sprintf("Установлен уровень логирования: %s", levelName)
	l.logMessage(ColorYellow, "INFO", message)
}

// функция для вывода логов с цветом
func (l *Logger) logMessage(color string, level string, message string) {
	l.logger.SetPrefix(fmt.Sprintf("%s[%s] %s", color, level, ColorReset))
	l.logger.Println(message)
}

func (l *Logger) Debug(message string) {
	if l.level == DEBUG {
		l.logMessage(ColorGreen, "DEBUG", message)
	}
}

func (l *Logger) Info(message string) {
	if l.level <= INFO {
		l.logMessage(ColorYellow, "INFO", message)
	}
}

func (l *Logger) Error(message string) {
	if l.level <= ERROR {
		l.logMessage(ColorRed, "ERROR", message)
	}
}
