package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)
//go:generate mockgen -source=main.go -destination=mock_interfaces.go -package=mocks
const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

func UnmarshalCandlesHistory(data []byte) (CandlesHistory, error) {
	var r CandlesHistory
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *CandlesHistory) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Currencies []string

func UnmarshalCurrencies(data []byte) (Currencies, error) {
	var r Currencies
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *Currencies) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOrderBook(data []byte) (OrderBook, error) {
	var r OrderBook
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *OrderBook) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OrderBook map[string]OrderBookPair
type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

type Ticker map[string]TickerValue

func UnmarshalTicker(data []byte) (Ticker, error) {
	var r Ticker
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *Ticker) Marshal() ([]byte, error) {
	return json.Marshal(r)
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
	Updated   int64  `json:"updated"`
}

func UnmarshalTrades(data []byte) (Trades, error) {
	var r Trades
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *Trades) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Trades map[string][]Pair

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     Type   `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type Type string

const (
	Buy  Type = "buy"
	Sell Type = "sell"
)

type GeneralIndicatorer interface {
	GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error)
}
type Indicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

// Функция для расчета простого скользящего среднего (SMA)
func calculateSMA(data []float64, period int) []float64 {
	var sma = make([]float64, len(data)/period)
	for i := range sma {
		sum := 0.0
		for _, d := range data[i*period : i*period+period] {
			sum += d
		}
		sma[i] = sum / float64(period)
	}
	return sma
}

// Функция для расчета экспоненциального скользящего среднего (EMA)
func calculateEMA(data []float64, period int) []float64 {
	if len(data) == 0 || period <= 0 {
		return nil
	}
	alpha := 2.0 / (float64(period) + 1.0)
	ema := make([]float64, len(data))
	ema[0] = data[0]
	for i := 1; i < len(data); i++ {
		ema[i] = alpha*data[i] + (1-alpha)*ema[i-1]
	}
	return ema
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, resolution int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, resolution int, start, end time.Time) ([]float64, error)
}

type GeneralIndicator struct{}

func (gi GeneralIndicator) GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, period, 3, from, to)
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo() *Exmo {
	exmo := &Exmo{}
	exmo.client = &http.Client{}
	exmo.url = "https://api.exmo.me/v1.1"
	return exmo

}

func (e *Exmo) GetTicker() (Ticker, error) {
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

	return UnmarshalTicker(body)
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {
	url := e.url + trades
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

		trades, err := UnmarshalTrades(body)
		if err != nil {
			return nil, err
		}

		for k, v := range trades {
			result[k] = v
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

		orderBook, err := UnmarshalOrderBook(body)
		if err != nil {
			return nil, err
		}

		for k, v := range orderBook {
			result[k] = v
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

	return UnmarshalCurrencies(body)

}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {

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

	return UnmarshalCandlesHistory(body)
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
		res = append(res, v.C)
	}

	return res, nil
}

type IndicatorSMA struct {
	exchanger Exchanger
}

func NewIndicatorSMA(exchange Exchanger) IndicatorSMA {
	return IndicatorSMA{
		exchanger: exchange,
	}
}
func (i IndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	data, err := i.exchanger.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}
	return calculateSMA(data, period), nil
}

type IndicatorEMA struct {
	exchanger Exchanger
}

func NewIndicatorEMA(exchange Exchanger) IndicatorEMA {
	return IndicatorEMA{
		exchanger: exchange,
	}
}

func (i IndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	data, err := i.exchanger.GetClosePrice(pair, limit, from, to)
	if err != nil {
		return nil, err
	}
	return calculateEMA(data, period), nil
}

func main() {
	var exchange Exchanger
	exchange = NewExmo()
	indicatorSMA := NewIndicatorSMA(exchange)
	generalIndicator := &GeneralIndicator{}
	sma, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorSMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sma)
	indicatorEMA := NewIndicatorEMA(exchange)
	ema, err := generalIndicator.GetData("BTC_USD", 30, time.Now().Add(-time.Hour*24*5), time.Now(), indicatorEMA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ema)

}
