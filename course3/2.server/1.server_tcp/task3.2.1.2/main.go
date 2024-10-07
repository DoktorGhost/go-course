package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	parts := strings.Fields(requestLine)
	if len(parts) < 3 {
		return
	}
	method := parts[0]
	path := parts[1]

	if method == "GET" && path == "/" {
		// Читаем содержимое index.html
		content, err := os.ReadFile("index.html")
		if err != nil {
			response := "HTTP/1.1 500 Internal Server Error\n" +
				"Content-Type: text/plain\n\n" +
				"500 Internal Server Error"
			conn.Write([]byte(response))
			return
		}

		response := "HTTP/1.1 200 OK\n" +
			"Content-Type: text/html\n\n" +
			string(content)
		conn.Write([]byte(response))
	} else {
		// Обрабатываем все остальные запросы
		response := "HTTP/1.1 404 Not Found\n" +
			"Content-Type: text/plain\n\n" +
			"404 Not Found"
		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}
		go handleConnection(conn)
	}
}
