package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	// Подключаемся к RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rate_limit_exceeded", // Очередь для сообщений о превышении лимита
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Получаем сообщения из очереди
	msgs, err := ch.Consume(
		q.Name, // очередь
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	// Обрабатываем сообщения
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// Отправляем email пользователю
			sendEmail(string(d.Body))
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Функция отправки email (упрощенная версия)
func sendEmail(message string) {
	fmt.Println("Message:", message)
}
