package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) Exchanger {
	var exmo Exmo

	if len(opts) == 0 {
		exmo.client = &http.Client{}
		exmo.url = "https://api.exmo.com/v1.1"
	}

	// Применение дополнительных опций, если они заданы
	for _, opt := range opts {
		opt(&exmo)
	}

	return &exmo
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
