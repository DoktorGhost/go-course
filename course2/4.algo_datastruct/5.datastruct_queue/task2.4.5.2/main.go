package main

import "fmt"

type BrowserHistory struct {
	stack []string
}

func (h *BrowserHistory) Visit(url string) {
	h.stack = append(h.stack, url)
	fmt.Println("Посещение URL-адреса:", url)
}

func (h *BrowserHistory) Back() {
	if len(h.stack) == 0 || len(h.stack) == 1 {
		fmt.Println("Нет больше истории для возврата")
		h.stack = make([]string, 0)
		return
	}

	h.stack = h.stack[:(len(h.stack) - 1)]
	index := len(h.stack) - 1
	element := h.stack[index]
	fmt.Println("Возврат к URL-адресу:", element)

}
func (h *BrowserHistory) PrintHistory() {
	if len(h.stack) == 0 {
		fmt.Println("История браузера пуста")
		return
	}
	fmt.Println("История браузера:")
	idx := len(h.stack) - 1
	for i := idx; i >= 0; i-- {
		fmt.Println(h.stack[i])
	}
}

func main() {
	history := &BrowserHistory{}
	history.Visit("www.google.com")
	history.Back()
	history.Visit("www.google.com")
	history.Visit("www.github.com")
	history.Visit("www.openai.com")
	history.Back()
	history.PrintHistory()
	history.Back()
	history.Back()
	history.Back()
	history.PrintHistory()
}
