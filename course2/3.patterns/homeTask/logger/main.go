package main

import (
	"student.vkusvill.ru/samsonov/go-course/course2/3.patterns/homeTask/logger/logger"
	"sync"
)

func main() {
	var log *logger.Logger
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			log = logger.GetInstance()
			defer wg.Done()
		}()
	}
	wg.Wait()

	log.SetLevel(logger.ERROR)

	log.Debug("Это сообщение отладки (DEBUG) (СООБЩЕНИЕ НЕ БУДЕТ ВЫВЕДЕНО)")
	log.Info("Это информационное сообщение (INFO) (СООБЩЕНИЕ НЕ БУДЕТ ВЫВЕДЕНО)")
	log.Error("Это сообщение ошибки (ERROR)")

	// меняем уровень логгирования
	log.SetLevel(logger.INFO)
	log.Debug("Это сообщение отладки (DEBUG) (СООБЩЕНИЕ НЕ БУДЕТ ВЫВЕДЕНО)")
	log.Info("Это информационное сообщение (INFO)")
	log.Error("Это сообщение ошибки (ERROR)")

	// меняем уровень логгирования
	log.SetLevel(logger.DEBUG)
	log.Debug("Это сообщение отладки (DEBUG)")
	log.Info("Это информационное сообщение (INFO)")
	log.Error("Это сообщение ошибки (ERROR)")
}
