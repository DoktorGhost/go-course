package main

import (
	"fmt"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type Indicatorer interface {
	SMA(period int) ([]float64, error)
	EMA(period int) ([]float64, error)
	LoadCandles(candles CandlesHistory)
}

type Indicator struct {
	candles CandlesHistory
}

func NewIndicator() *Indicator {
	return &Indicator{}
}

func (i *Indicator) SMA(period int) ([]float64, error) {
	// Рассчитать SMA на основе данных свечей
	return []float64{}, nil
}

func (i *Indicator) EMA(period int) ([]float64, error) {
	// Рассчитать EMA на основе данных свечей
	return []float64{}, nil
}

func (i *Indicator) LoadCandles(candles CandlesHistory) {
	i.candles = candles
}

// Dashboarder must return candles history and indicators with several periods, presets by opts
type Dashboarder interface {
	GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     map[string][]IndicatorData
	Period         int
	From           time.Time
	To             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

type Dashboard struct {
	exchange           Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	Period             int
	From               time.Time
	To                 time.Time
}

func NewDashboard(exchange Exchanger) *Dashboard {
	return &Dashboard{exchange: exchange}
}

func (d *Dashboard) GetDashboard(pair string, opts ...func(*Dashboard)) (DashboardData, error) {
	// Применение опций
	for _, opt := range opts {
		opt(d)
	}

	// Получение истории свечей
	var candlesHistory CandlesHistory
	if d.withCandlesHistory {
		var err error
		candlesHistory, err = d.exchange.GetCandlesHistory(pair, d.Period, d.From, d.To)
		if err != nil {
			return DashboardData{}, err
		}
	}

	// Подготовка индикаторов
	indicators := make(map[string][]IndicatorData)
	for _, opt := range d.IndicatorOpts {
		opt.Indicator.LoadCandles(candlesHistory)
		for _, period := range opt.Periods {
			var values []float64
			var err error
			switch opt.Name {
			case "SMA":
				values, err = opt.Indicator.SMA(period)
			case "EMA":
				values, err = opt.Indicator.EMA(period)
			}
			if err != nil {
				return DashboardData{}, err
			}
			indicators[opt.Name] = append(indicators[opt.Name], IndicatorData{
				Name:     opt.Name,
				Period:   period,
				Indicate: values,
			})
		}
	}

	return DashboardData{
		Name:           pair,
		CandlesHistory: candlesHistory,
		Indicators:     indicators,
		Period:         d.Period,
		From:           d.From,
		To:             d.To,
	}, nil
}

func WithCandlesHistory(period int, from, to time.Time) func(*Dashboard) {
	// Реализация для задания истории свечей
	return func(d *Dashboard) {
		d.withCandlesHistory = true
		d.Period = period
		d.From = from
		d.To = to
	}
}

func WithIndicatorOpts(opts ...IndicatorOpt) func(*Dashboard) {
	// Реализация для задания индикаторов
	return func(d *Dashboard) {
		d.IndicatorOpts = append(d.IndicatorOpts, opts...)
	}
}

func main() {
	exchange := NewExmo()
	dashboard := NewDashboard(exchange)

	data, err := dashboard.GetDashboard(
		"BTC_USD",
		WithCandlesHistory(30, time.Now().Add(-time.Hour*3), time.Now()),
		WithIndicatorOpts(
			IndicatorOpt{
				Name:      "SMA",
				Periods:   []int{5, 10, 20},
				Indicator: NewIndicator(),
			},
			IndicatorOpt{
				Name:      "EMA",
				Periods:   []int{5, 10, 20},
				Indicator: NewIndicator(),
			},
		),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

/*
Создан дашборд фасад.
Обертка предоставляет интерфейс Dashboarder.
Фасад скрывает сложность работы с данными о свечах и индикаторах.
Метод GetDashboard возвращает данные о свечах и индикаторах.
Метод GetDashboard принимает параметры для настройки запроса.
Метод GetDashboard возвращает ожидаемый результат.
Решение расположи по следующему пути: course2/3.pa erns/4.pa erns_facade/task2.3.4.1

*/
