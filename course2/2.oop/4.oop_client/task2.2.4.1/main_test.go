package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewExmoDefault(t *testing.T) {
	// Создаем новый объект Exmo без опций
	exmo := NewExmo()

	// Проверяем, что client установлен на *http.Client{}
	if exmo.client == nil {
		t.Error("expected client to be initialized, got nil")
	}

	// Проверяем, что URL установлен на "https://api.exmo.me/v1.1"
	expectedURL := "https://api.exmo.me/v1.1"
	if exmo.url != expectedURL {
		t.Errorf("expected URL %s, got %s", expectedURL, exmo.url)
	}
}

func mockTickerHandler(w http.ResponseWriter, r *http.Request) {
	tickerResponse := Ticker{
		"BTC_USD": TickerValue{
			BuyPrice:  "50000",
			SellPrice: "51000",
			LastTrade: "50500",
			High:      "52000",
			Low:       "49000",
			Avg:       "50500",
			Vol:       "1000",
			VolCurr:   "10000000",
			Updated:   1633065600,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickerResponse)
}

func TestGetTicker(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockTickerHandler))
	defer ts.Close()

	exmo := NewExmo(WithClient(ts.Client()), WithURL(ts.URL))
	tickers, err := exmo.GetTicker()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedTicker := Ticker{
		"BTC_USD": TickerValue{
			BuyPrice:  "50000",
			SellPrice: "51000",
			LastTrade: "50500",
			High:      "52000",
			Low:       "49000",
			Avg:       "50500",
			Vol:       "1000",
			VolCurr:   "10000000",
			Updated:   1633065600,
		},
	}

	if !equalTicker(tickers, expectedTicker) {
		t.Errorf("expected %v, got %v", expectedTicker, tickers)
	}
}

func equalTicker(a, b Ticker) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bVal, ok := b[k]; !ok || v != bVal {
			return false
		}
	}
	return true
}

func mockTradesHandler(w http.ResponseWriter, r *http.Request) {
	tradesResponse := Trades{
		"BTC_USD": []Pair{
			{
				TradeID:  1,
				Type:     "buy",
				Price:    "50000",
				Quantity: "0.1",
				Amount:   "5000",
				Date:     1633065600,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tradesResponse)
}

func TestGetTrades(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockTradesHandler))
	defer ts.Close()

	exmo := NewExmo(WithClient(ts.Client()), WithURL(ts.URL))
	trades, err := exmo.GetTrades("BTC_USD")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedTrades := Trades{
		"BTC_USD": []Pair{
			{
				TradeID:  1,
				Type:     "buy",
				Price:    "50000",
				Quantity: "0.1",
				Amount:   "5000",
				Date:     1633065600,
			},
		},
	}

	if !equalTrades(trades, expectedTrades) {
		t.Errorf("expected %v, got %v", expectedTrades, trades)
	}
}

func equalTrades(a, b Trades) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bVal, ok := b[k]; !ok || !equalPairs(v, bVal) {
			return false
		}
	}
	return true
}

func equalPairs(a, b []Pair) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Mock server response for OrderBook
func mockOrderBookHandler(w http.ResponseWriter, r *http.Request) {
	orderBookResponse := OrderBook{
		"BTC_USD": OrderBookPair{
			AskQuantity: "0.1",
			AskAmount:   "1000",
			AskTop:      "50000",
			BigQuantity: "0.2",
			BigAmount:   "2000",
			BigTop:      "50500",
			Ask: [][]string{
				{"50000", "0.1"},
				{"50010", "0.2"},
			},
			Bid: [][]string{
				{"49500", "0.1"},
				{"49490", "0.2"},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderBookResponse)
}

func TestGetOrderBook(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockOrderBookHandler))
	defer ts.Close()

	exmo := NewExmo(WithClient(ts.Client()), WithURL(ts.URL))
	orderBook, err := exmo.GetOrderBook(10, "BTC_USD")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedOrderBook := OrderBook{
		"BTC_USD": OrderBookPair{
			AskQuantity: "0.1",
			AskAmount:   "1000",
			AskTop:      "50000",
			BigQuantity: "0.2",
			BigAmount:   "2000",
			BigTop:      "50500",
			Ask: [][]string{
				{"50000", "0.1"},
				{"50010", "0.2"},
			},
			Bid: [][]string{
				{"49500", "0.1"},
				{"49490", "0.2"},
			},
		},
	}

	if !equalOrderBook(orderBook, expectedOrderBook) {
		t.Errorf("expected %v, got %v", expectedOrderBook, orderBook)
	}
}

func equalOrderBook(a, b OrderBook) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bVal, ok := b[k]; !ok || !equalOrderBookPair(v, bVal) {
			return false
		}
	}
	return true
}

func equalOrderBookPair(a, b OrderBookPair) bool {
	if a.AskQuantity != b.AskQuantity || a.AskAmount != b.AskAmount || a.AskTop != b.AskTop ||
		a.BigQuantity != b.BigQuantity || a.BigAmount != b.BigAmount || a.BigTop != b.BigTop {
		return false
	}
	if !equal2DStringSlice(a.Ask, b.Ask) || !equal2DStringSlice(a.Bid, b.Bid) {
		return false
	}
	return true
}

func equal2DStringSlice(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// Mock server response for Currencies
func mockCurrenciesHandler(w http.ResponseWriter, r *http.Request) {
	currenciesResponse := Currencies{"USDT", "RUB", "BTC", "DOGE", "LTC"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currenciesResponse)
}

func TestGetCurrencies(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockCurrenciesHandler))
	defer ts.Close()

	exmo := NewExmo(WithClient(ts.Client()), WithURL(ts.URL))
	currencies, err := exmo.GetCurrencies()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedCurrencies := Currencies{"USDT", "RUB", "BTC", "DOGE", "LTC"}

	if !equalCurrencies(currencies, expectedCurrencies) {
		t.Errorf("expected %v, got %v", expectedCurrencies, currencies)
	}
}

func equalCurrencies(a, b Currencies) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Mock server response for CandlesHistory
func mockCandlesHistoryHandler(w http.ResponseWriter, r *http.Request) {
	candlesResponse := CandlesHistory{
		Candles: []Candle{
			{
				Time:   1585557000000,
				Open:   6590.6164,
				Close:  6602.3624,
				High:   6618.78965693,
				Low:    6579.054,
				Volume: 6.932754980000013,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candlesResponse)
}

func TestGetCandlesHistory(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockCandlesHistoryHandler))
	defer ts.Close()

	exmo := NewExmo(WithClient(ts.Client()), WithURL(ts.URL))
	candlesHistory, err := exmo.GetCandlesHistory("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedCandlesHistory := CandlesHistory{
		Candles: []Candle{
			{
				Time:   1585557000000,
				Open:   6590.6164,
				Close:  6602.3624,
				High:   6618.78965693,
				Low:    6579.054,
				Volume: 6.932754980000013,
			},
		},
	}

	if !equalCandlesHistory(candlesHistory, expectedCandlesHistory) {
		t.Errorf("expected %v, got %v", expectedCandlesHistory, candlesHistory)
	}
}

func equalCandlesHistory(a, b CandlesHistory) bool {
	if len(a.Candles) != len(b.Candles) {
		return false
	}
	for i := range a.Candles {
		if a.Candles[i] != b.Candles[i] {
			return false
		}
	}
	return true
}

// Mock server response for ClosePrice
func mockClosePriceHandler(w http.ResponseWriter, r *http.Request) {
	candlesResponse := CandlesHistory{
		Candles: []Candle{
			{
				Time:   1585557000000,
				Open:   6590.6164,
				Close:  6602.3624,
				High:   6618.78965693,
				Low:    6579.054,
				Volume: 6.932754980000013,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candlesResponse)
}

func TestGetClosePrice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(mockClosePriceHandler))
	defer ts.Close()

	exmo := NewExmo(WithClient(ts.Client()), WithURL(ts.URL))
	closePrices, err := exmo.GetClosePrice("BTC_USD", 30, time.Now().Add(-time.Hour*24), time.Now())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedClosePrices := []float64{6602.3624}

	if !equalFloat64Slice(closePrices, expectedClosePrices) {
		t.Errorf("expected %v, got %v", expectedClosePrices, closePrices)
	}
}

func equalFloat64Slice(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
