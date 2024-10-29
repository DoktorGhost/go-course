package broker

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func sendToRabbitMQ(message string) {
	conn, err := amqp.Dial("amqp://" + os.Getenv("RABBITMQ_USER") + ":" + os.Getenv("RABBITMQ_PASS") + "@" + os.Getenv("RABBITMQ_HOST") + ":" + os.Getenv("RABBITMQ_PORT") + "/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rate_limit_exceeded",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}

	log.Println("Sent message to RabbitMQ:", message)
}

func sendToKafka(message string) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:    os.Getenv("KAFKA_TOPIC"),
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(message),
	})
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	} else {
		log.Println("Sent message to Kafka:", message)
	}
}

func SendMessage(message string) {
	brokerType := os.Getenv("MESSAGE_BROKER")

	switch brokerType {
	case "kafka":
		sendToKafka(message)
	case "rabbitmq":
		sendToRabbitMQ(message)
	default:
		log.Println("Unknown broker type:", brokerType)
	}
}
