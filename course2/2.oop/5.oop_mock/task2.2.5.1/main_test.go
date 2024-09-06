package main

import (
	"github.com/golang/mock/gomock"
	"math"
	"reflect"
	"testing"
	"time"
)

func roundSlice(slice []float64, precision int) []float64 {
	var factor float64 = math.Pow(10, float64(precision))
	for i, v := range slice {
		slice[i] = math.Round(v*factor) / factor
	}
	return slice
}

func TestIndicator_SMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExchanger := NewMockExchanger(ctrl)
	mockExchanger.EXPECT().
		GetClosePrice("BTC_USD", 30, gomock.Any(), gomock.Any()).
		Return([]float64{22.27, 22.19, 22.08, 22.17, 22.18}, nil).
		Times(1)

	indicator := NewIndicator(mockExchanger, WithCalculateSMA(calculateSMA))

	sma, err := indicator.SMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedSMA := []float64{22.178}
	// Округление до 3-х десятичных знаков
	sma = roundSlice(sma, 3)
	expectedSMA = roundSlice(expectedSMA, 3)

	if !reflect.DeepEqual(sma, expectedSMA) {
		t.Fatalf("Expected %v, but got %v", expectedSMA, sma)
	}
}

func TestIndicator_EMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExchanger := NewMockExchanger(ctrl)
	mockExchanger.EXPECT().
		GetClosePrice("BTC_USD", 30, gomock.Any(), gomock.Any()).
		Return([]float64{22.27, 22.19, 22.08, 22.17, 22.18}, nil).
		Times(1)

	indicator := NewIndicator(mockExchanger, WithCalculateEMA(calculateEMA))

	ema, err := indicator.EMA("BTC_USD", 30, 5, time.Now().AddDate(0, 0, -2), time.Now())
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedEMA := []float64{22.178}
	// Округление до 3-х десятичных знаков
	ema = roundSlice(ema, 3)
	expectedEMA = roundSlice(expectedEMA, 3)

	if !reflect.DeepEqual(ema, expectedEMA) {
		t.Fatalf("Expected %v, but got %v", expectedEMA, ema)
	}
}
