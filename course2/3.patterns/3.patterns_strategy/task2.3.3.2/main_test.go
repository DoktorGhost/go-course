package main

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIndicatorSMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок Exchanger
	mockExchanger := NewMockExchanger(ctrl)

	// Ожидаем вызова GetClosePrice с любым значением для пары и проверяем результат
	expectedData := []float64{10, 10, 11, 12, 15}
	mockExchanger.EXPECT().GetClosePrice("BTC_USD", 30, gomock.Any(), gomock.Any()).Return(expectedData, nil)

	// Создаем экземпляр IndicatorSMA
	indicatorSMA := NewIndicatorSMA(mockExchanger)
	generalIndicator := &GeneralIndicator{}

	// Получаем результат из GeneralIndicator
	sma, err := generalIndicator.GetData("BTC_USD", 30, time.Now(), time.Now(), indicatorSMA)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Рассчитываем ожидаемый результат SMA
	expectedSMA := calculateSMA(expectedData, 3)

	// Проверяем, что полученные данные соответствуют ожидаемым
	assert.ElementsMatch(t, expectedSMA, sma, "SMA values do not match")
}

func TestIndicatorEMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок Exchanger
	mockExchanger := NewMockExchanger(ctrl)

	// Ожидаем вызова GetClosePrice с любым значением для пары и проверяем результат
	expectedData := []float64{10, 10, 11, 12, 15}
	mockExchanger.EXPECT().GetClosePrice("BTC_USD", 30, gomock.Any(), gomock.Any()).Return(expectedData, nil)

	// Создаем экземпляр IndicatorSMA
	indicatorEMA := NewIndicatorEMA(mockExchanger)
	generalIndicator := &GeneralIndicator{}

	// Получаем результат из GeneralIndicator
	ema, err := generalIndicator.GetData("BTC_USD", 30, time.Now(), time.Now(), indicatorEMA)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Рассчитываем ожидаемый результат SMA
	expectedEMA := calculateEMA(expectedData, 3)

	// Проверяем, что полученные данные соответствуют ожидаемым
	assert.ElementsMatch(t, expectedEMA, ema, "SMA values do not match")
}
