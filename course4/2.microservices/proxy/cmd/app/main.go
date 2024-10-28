package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/streadway/amqp"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var ipLimits sync.Map

type IPLimiter struct {
	Limiter *limiter
}

func main() {
	loadEnv()
	http.HandleFunc("/api/user/", proxyHandler("http://user:8080"))
	http.HandleFunc("/api/auth/", proxyHandler("http://auth:8080"))
	http.HandleFunc("/api/address/", rateLimitMiddleware(proxyHandler("http://geo:8080")))

	log.Println("Proxy server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func proxyHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxyURL, err := url.Parse(target)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}

		r.URL.Scheme = proxyURL.Scheme
		r.URL.Host = proxyURL.Host

		resp, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)
		limiter := getLimiterForIP(clientIP)

		if !limiter.Limiter.Take() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("429 Too Many Requests - rate limit exceeded"))
			sendMessage("Rate limit exceeded for IP: " + clientIP)
			return
		}

		next(w, r)
	}
}

func getClientIP(r *http.Request) string {
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		return forwardedFor
	}
	return r.RemoteAddr
}

func getLimiterForIP(ip string) *IPLimiter {
	value, exists := ipLimits.Load(ip)
	if exists {
		return value.(*IPLimiter)
	}

	limiter := &IPLimiter{
		Limiter: newLimiter(5, time.Minute),
	}
	ipLimits.Store(ip, limiter)
	return limiter
}

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

func sendMessage(message string) {
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

type limiter struct {
	count  int
	start  time.Time
	limit  int
	period time.Duration
}

func newLimiter(limit int, period time.Duration) *limiter {
	startTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	return &limiter{0, startTime, limit, period}
}

func (l *limiter) Take() bool {
	currentTime := time.Now()

	if l.count == 0 {
		l.count++
		l.start = currentTime
		return true
	}

	if l.count < l.limit {
		l.count++
		return true
	}

	if currentTime.Sub(l.start) < l.period {
		return false
	}

	l.count = 1
	l.start = currentTime
	return true
}
