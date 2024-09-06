package main

import (
	"encoding/json"
	"fmt"
	"github.com/cinar/indicator/v2/helper"
	"github.com/cinar/indicator/v2/trend"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	Time   int64   `json:"t"`
	Open   float64 `json:"o"`
	Close  float64 `json:"c"`
	High   float64 `json:"h"`
	Low    float64 `json:"l"`
	Volume float64 `json:"v"`
}

type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BigQuantity string     `json:"bid_quantity"`
	BigAmount   string     `json:"bid_amount"`
	BigTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}
type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int    `json:"updated"`
}

type Pair struct {
	TradeID  int    `json:"trade_id"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Amount   string `json:"amount"`
	Date     int64  `json:"date"`
}

type Currencies []string
type OrderBook map[string]OrderBookPair
type Ticker map[string]TickerValue
type Trades map[string][]Pair

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

//go:generate mockgen -source=main.go -destination=mock_client.go -package=main
type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{}
	if len(opts) == 0 {
		exmo.client = &http.Client{}
		exmo.url = "https://api.exmo.me/v1.1"
		return exmo
	}
	for _, opt := range opts {
		opt(exmo)
	}

	return exmo
}
func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}
func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
	}
}

func (e *Exmo) GetTicker() (Ticker, error) {
	var result Ticker

	url := e.url + ticker

	method := "POST"
	payload := strings.NewReader("")
	client := e.client
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	url := fmt.Sprintf("%s%s", e.url, trades)
	method := "POST"
	client := e.client
	var result Trades

	for _, pair := range pairs {
		payload := strings.NewReader("pair=" + pair)
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}
	}

	return result, nil
}
func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	url := e.url + orderBook
	method := "POST"
	client := e.client
	var result OrderBook

	for _, pair := range pairs {
		payload := strings.NewReader("pair=" + pair + "&limit=" + strconv.Itoa(limit))
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}
	}

	return result, nil
}
func (e *Exmo) GetCurrencies() (Currencies, error) {
	url := e.url + currency
	method := "POST"

	payload := strings.NewReader("")

	client := e.client
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var currencies Currencies
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return currencies, nil

}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	var result CandlesHistory
	url := fmt.Sprintf("%s%s?symbol=%s&resolution=%d&from=%v&to=%v", e.url, candlesHistory, pair, limit, start.Unix(), end.Unix())

	method := "GET"
	client := e.client
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return CandlesHistory{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return CandlesHistory{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return CandlesHistory{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return CandlesHistory{}, err
	}

	return result, nil
}
func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	var result CandlesHistory
	url := fmt.Sprintf("%s%s?symbol=%s&resolution=%d&from=%v&to=%v", e.url, candlesHistory, pair, limit, start.Unix(), end.Unix())

	method := "GET"
	client := e.client
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	var res []float64

	for _, v := range result.Candles {
		res = append(res, v.Close)
	}

	return res, nil
}

func main() {
	var exchange Exchanger
	exchange = NewExmo()
	ind := NewIndicator(exchange, WithCalculateEMA(calculateEMA), WithCalculateSMA(calculateSMA))
	sma, err := ind.SMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sma)
	ema, err := ind.EMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ema)
}

type Indicatorer interface {
	SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

type IndicatorOption func(*Indicator)

func NewIndicator(exchange Exchanger, options ...IndicatorOption) *Indicator {
	ind := &Indicator{
		exchange: exchange,
	}
	for _, option := range options {
		option(ind)
	}
	return ind
}

func WithCalculateEMA(calculateEMA func(data []float64, period int) []float64) IndicatorOption {
	return func(ind *Indicator) {
		ind.calculateEMA = calculateEMA
	}
}
func WithCalculateSMA(calculateSMA func(data []float64, period int) []float64) IndicatorOption {
	return func(ind *Indicator) {
		ind.calculateSMA = calculateSMA
	}
}

// SMA - метод для расчета простого скользящего среднего
func (ind *Indicator) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	data, err := ind.exchange.GetClosePrice(pair, resolution, from, to)
	if err != nil {
		return nil, err
	}
	return ind.calculateSMA(data, period), nil
}

// EMA - метод для расчета экспоненциального скользящего среднего
func (ind *Indicator) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	data, err := ind.exchange.GetClosePrice(pair, resolution, from, to)
	if err != nil {
		return nil, err
	}
	return ind.calculateEMA(data, period), nil
}

// Функция для расчета простого скользящего среднего (SMA)
func calculateSMA(closing []float64, period int) []float64 {
	// Преобразуем срез closing в канал
	input := helper.SliceToChan(closing)

	// Создаем объект SMA с указанным периодом
	sma := trend.NewSmaWithPeriod[float64](period)

	// Рассчитываем SMA
	output := sma.Compute(input)

	// Преобразуем канал обратно в срез и возвращаем результат
	return helper.ChanToSlice(output)
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func calculateEMA(closing []float64, period int) []float64 {
	// Преобразуем срез closing в канал
	input := helper.SliceToChan(closing)

	// Создаем объект SMA с указанным периодом
	ema := trend.NewEmaWithPeriod[float64](period)

	// Рассчитываем SMA
	output := ema.Compute(input)

	// Преобразуем канал обратно в срез и возвращаем результат
	return helper.ChanToSlice(output)
}
