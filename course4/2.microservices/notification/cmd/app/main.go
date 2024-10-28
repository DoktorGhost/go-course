package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

// Устанавливаем флаг, чтобы знать, откуда пришло сообщение
const sourceRabbitMQ = "rabbitmq"
const sourceKafka = "kafka"

// Конфигурация для RabbitMQ
const rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
const rabbitMQQueue = "rate_limit_exceeded"

// Конфигурация для Kafka
const kafkaBroker = "kafka:9092"
const kafkaTopic = "rate_limit_exceeded"

func main() {
	var wg sync.WaitGroup

	// Запускаем консьюмер для RabbitMQ
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumeFromRabbitMQ()
	}()

	// Запускаем консьюмер для Kafka
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumeFromKafka()
	}()

	log.Println("Listening for messages. To exit press CTRL+C")
	wg.Wait()
}

// Функция для получения сообщений из RabbitMQ
func consumeFromRabbitMQ() {
	// Подключаемся к RabbitMQ
	var conn *amqp.Connection
	var err error
	for {
		conn, err = amqp.Dial(rabbitMQURL)
		if err != nil {
			log.Println(err)
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rabbitMQQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	// Получаем сообщения из очереди
	msgs, err := ch.Consume(
		q.Name,
		"",
		true, // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	for d := range msgs {
		log.Printf("Received a message from RabbitMQ: %s", d.Body)
		processMessage(string(d.Body), sourceRabbitMQ)
	}
}

// Функция для получения сообщений из Kafka
func consumeFromKafka() {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   kafkaTopic,
		GroupID: "email_service",
	})
	defer kafkaReader.Close()

	for {
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message from Kafka: %v", err)
		}
		log.Printf("Received a message from Kafka: %s", string(msg.Value))
		processMessage(string(msg.Value), sourceKafka)
	}
}

// Функция для обработки сообщения
func processMessage(message string, source string) {
	fmt.Printf("Processing message from %s: %s\n", source, message)
	sendEmail(message) // Например, отправляем email
}

// Функция отправки email (упрощенная версия)
func sendEmail(message string) {
	fmt.Println("Message:", message)
}
