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

type _logger struct {
	logger *log.Logger
	level  int
}

var (
	once           sync.Once
	singleInstance *_logger
)

func Init() {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = &_logger{
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

}

func SetLevel(level int) {
	singleInstance.level = level
	levelName := "INFO"
	if level == 0 {
		levelName = "DEBUG"
	} else if level == 2 {
		levelName = "ERROR"
	}
	message := fmt.Sprintf("Установлен уровень логирования: %s", levelName)
	singleInstance.logMessage(ColorYellow, "INFO", message)
}

// функция для вывода логов с цветом
func (l *_logger) logMessage(color string, level string, message string) {
	l.logger.SetPrefix(fmt.Sprintf("%s[%s] %s", color, level, ColorReset))
	l.logger.Println(message)
}

func Debug(message string) {
	if singleInstance.level == DEBUG {
		singleInstance.logMessage(ColorGreen, "DEBUG", message)
	}
}

func Info(message string) {
	if singleInstance.level <= INFO {
		singleInstance.logMessage(ColorYellow, "INFO", message)
	}
}

func Error(message string) {
	if singleInstance.level <= ERROR {
		singleInstance.logMessage(ColorRed, "ERROR", message)
	}
}
