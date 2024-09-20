package main

import (
	"bufio"
	"fmt"
	"github.com/DoktorGhost/randomanimals"
	"net"
	"os"
	"strings"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

// будем окрашивать информационные сообщения
const (
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
)

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Сервер запущен. Ожидание подключений...")

	go broadcaster()

	// Принимаем входящие подключения
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	ch := make(chan string)
	//go clientWriter(conn, ch)

	//who := conn.RemoteAddr().String()
	who := randomanimals.RandomAnimal()
	go clientWriter(conn, ch, who)

	cli := client{conn, who, ch}

	ch <- colorText(ColorYellow, "Ваш ник: "+who)
	messages <- colorText(ColorGreen, who+" подключился к нам")
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- colorText(ColorRed, who+" покинул нас")

	conn.Close()
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string, who string) {
	for msg := range ch {
		if strings.HasPrefix(msg, who) {
			// Сообщение от текущего клиента
			msg = colorText(ColorBlue, msg)
		} else {
			// Сообщение от других клиентов
			msg = colorText(ColorPurple, msg)
		}
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			fmt.Println("Ошибка при отправке данных клиенту:", err)
			return
		}
	}
}

func colorText(colorCode string, text string) string {
	return fmt.Sprintf("%s%s%s", colorCode, text, "\033[0m")
}
