package main

import (
	"github.com/streadway/amqp"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Карта для хранения лимитеров по IP-адресам
var ipLimits sync.Map

// Структура для хранения лимитера
type IPLimiter struct {
	Limiter *limiter
}

// Главная функция сервера
func main() {
	// Настраиваем обработчики для каждого сервиса
	http.HandleFunc("/api/user/", proxyHandler("http://user:8080"))
	http.HandleFunc("/api/auth/", proxyHandler("http://auth:8080"))
	http.HandleFunc("/api/address/", rateLimitMiddleware(proxyHandler("http://geo:8080")))

	// Запускаем прокси-сервер на порту 8080
	log.Println("Proxy server started on :8080")
	http.ListenAndServe(":8080", nil)
}

// Функция для обработки проксированных запросов
func proxyHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Проксирование запроса к целевому сервису
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

// Middleware для проверки лимита запросов по IP
func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем IP адрес клиента
		clientIP := getClientIP(r)

		// Получаем или создаем лимитер для данного IP
		limiter := getLimiterForIP(clientIP)

		if !limiter.Limiter.Take() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("429 Too Many Requests - rate limit exceeded"))
			sendToRabbitMQ("Rate limit exceeded for IP: " + clientIP)
			return
		}

		// Если лимит не превышен, передаем управление следующему обработчику
		next(w, r)
	}
}

// Функция для извлечения IP адреса клиента
func getClientIP(r *http.Request) string {
	// Проверяем наличие заголовков X-Real-IP и X-Forwarded-For (если за прокси)
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		return forwardedFor
	}
	return r.RemoteAddr // используем IP адрес, если запрос не за прокси
}

func getLimiterForIP(ip string) *IPLimiter {
	// Проверяем, существует ли лимитер для данного IP
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

// Функция для отправки сообщений в RabbitMQ
func sendToRabbitMQ(message string) {
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
		"rate_limit_exceeded",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
}

//свой лимитер

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

	// Если это первый запрос
	if l.count == 0 {
		l.count++
		l.start = currentTime
		return true
	}

	// Если количество запросов меньше лимита
	if l.count < l.limit {
		l.count++
		return true
	}

	// Проверяем, прошел ли период времени с момента первого запроса
	if currentTime.Sub(l.start) < l.period {
		return false // Лимит превышен
	}

	// Сброс счетчика и установка нового времени старта
	l.count = 1
	l.start = currentTime
	return true
}
