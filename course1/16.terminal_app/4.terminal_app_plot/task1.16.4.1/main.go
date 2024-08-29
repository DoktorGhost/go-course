package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/guptarohit/asciigraph"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	width  = 100 // Ширина графика
	height = 10  // Высота графика
)

var (
	symbols       = []string{"BTC_USD", "LTC_USD", "ETH_USD"}
	menuDisplayed = true
	coin          = ""
	newData       = make([]float64, 0, width)
)

func main() {
	dataBTC := make([]float64, 0, width)
	dataLTC := make([]float64, 0, width)
	dataETH := make([]float64, 0, width)
	lastB := 0.0
	lastL := 0.0
	lastE := 0.0
	last := 0.0

	var muBTC sync.Mutex
	var muLTC sync.Mutex
	var muETH sync.Mutex

	go collectData("BTC_USD", &lastB, &dataBTC, &muBTC)
	go collectData("LTC_USD", &lastL, &dataLTC, &muLTC)
	go collectData("ETH_USD", &lastE, &dataETH, &muETH)

	go handleInput()

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	displayMenu()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {

		if menuDisplayed {
			continue
		}

		if coin == "BTC_USD" {
			muBTC.Lock()
			newData = dataBTC
			last = lastB
			muBTC.Unlock()
		} else if coin == "LTC_USD" {
			muLTC.Lock()
			newData = dataLTC
			last = lastL
			muLTC.Unlock()
		} else if coin == "ETH_USD" {
			muETH.Lock()
			newData = dataETH
			last = lastE
			muETH.Unlock()
		}

		caption := fmt.Sprintf("\nТекущий курс: %.2f %s\nДата: %s\nВремя: %s\n",
			last, coin, time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"))

		if len(newData) < 1 {
			continue
		}
		graph := asciigraph.Plot(newData,
			asciigraph.Height(height),
			asciigraph.Width(width),
			asciigraph.SeriesColors(asciigraph.Red),
		)

		clearScreen()
		fmt.Fprintln(writer, graph, caption)
		time.Sleep(time.Second * 1)

	}

}

// очистка консоли
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// menu
func displayMenu() {
	clearScreen()
	fmt.Println("Меню:")
	for i, symbol := range symbols {
		fmt.Printf("%d: %s\n", i+1, symbol)
	}
	fmt.Println("Press 'q' to quit or BACKSPACE to return to the menu")
}

// функция для считывания клавиатуры
func handleInput() {
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard:", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'q':
			os.Exit(0)
		case 'й':
			os.Exit(0)
		case '\x00':
			menuDisplayed = true
			displayMenu()
		default:
			if char >= '1' && char <= '3' {
				index := int(char - '1')
				if index >= 0 && index < len(symbols) {
					menuDisplayed = false
					coin = symbols[index]
				}
			}
		}

	}

}

// запрос на сайт
func requestAPI(coin string) (float64, error) {
	var url string
	switch coin {
	case "BTC_USD":
		url = "https://www.binance.com/ru/price/bitcoin"
	case "LTC_USD":
		url = "https://www.binance.com/ru/price/litecoin"
	case "ETH_USD":
		url = "https://www.binance.com/ru/price/ethereum"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	// Чтение всего содержимого ответа
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	// Вывод содержимого ответа
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return 0, err
	}

	// Извлечение текста цены
	priceText := doc.Find(".css-1bwgsh3").Text()

	// Очистка текста от нецифровых символов, кроме точки
	priceText = strings.TrimSpace(priceText)
	priceText = strings.ReplaceAll(priceText, "$", "") // Удаление символа доллара
	priceText = strings.ReplaceAll(priceText, ",", "")
	priceText = strings.TrimSpace(priceText)

	// Преобразование текста в float64
	price, err := strconv.ParseFloat(priceText, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

// создание слайса с курсом
func collectData(coin string, last *float64, dataBuffer *[]float64, mu *sync.Mutex) {
	for {
		newData, err := requestAPI(coin)
		if err != nil {
			fmt.Println("Error collecting data for", coin, ":", err)
			time.Sleep(1 * time.Second)
			continue
		}

		mu.Lock()
		if len(*dataBuffer) >= width {
			*dataBuffer = (*dataBuffer)[1:]
		}
		*dataBuffer = append(*dataBuffer, newData)
		*last = newData
		mu.Unlock()

		time.Sleep(1 * time.Second)
	}
}
