package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// подключиться к серверу
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		os.Exit(1)
	}

	// запустить горутину, которая будет читать все сообщения от сервера и выводить их в консоль
	go clientReader(conn)

	// читать сообщения от stdin и отправлять их на сервер
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		_, err = conn.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			fmt.Println("Ошибка при отправке данных:", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}
}

// clientReader выводит на экран все сообщения от сервера
func clientReader(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении данных:", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}
