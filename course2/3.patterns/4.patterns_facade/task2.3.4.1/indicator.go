package main

import (
	"github.com/cinar/indicator"
)

type TechnicalAnalysis interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA(periods ...int) [][]float64
}

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *Lines) StochPrice() ([]float64, []float64) {
	k, d := indicator.StochasticOscillator(t.high, t.low, t.closing)
	return k, d
}

func (t *Lines) RSI(period int) ([]float64, []float64) {
	rs, rsi := indicator.RsiPeriod(period, t.closing)
	return rs, rsi
}

func (t *Lines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	_, rsi := t.RSI(rsiPeriod)
	k, d := indicator.StochasticOscillator(rsi, rsi, rsi)
	return k, d
}

func Smooth(period int, values []float64) []float64 {
	return indicator.Sma(period, values)
}

func SmoothEMA(period int, values []float64) []float64 {
	return indicator.Ema(period, values)
}

func (t *Lines) SMA(period int) []float64 {
	return indicator.Sma(period, t.closing)
}

func (t *Lines) MACD() ([]float64, []float64) {
	return indicator.Macd(t.closing)
}

func (t *Lines) EMA(periods ...int) [][]float64 {
	var res [][]float64
	for i := range periods {
		res = append(res, indicator.Ema(periods[i], t.closing))
	}
	return res
}

func LoadCandles(candles CandlesHistory) *Lines {
	t := &Lines{}
	for i := range candles.Candles {
		t.closing = append(t.closing, candles.Candles[i].C)
		t.low = append(t.low, candles.Candles[i].L)
		t.high = append(t.high, candles.Candles[i].H)
	}
	return t
}
