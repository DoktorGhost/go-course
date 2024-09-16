// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockExchanger is a mock of Exchanger interface.
type MockExchanger struct {
	ctrl     *gomock.Controller
	recorder *MockExchangerMockRecorder
}

// MockExchangerMockRecorder is the mock recorder for MockExchanger.
type MockExchangerMockRecorder struct {
	mock *MockExchanger
}

// NewMockExchanger creates a new mock instance.
func NewMockExchanger(ctrl *gomock.Controller) *MockExchanger {
	mock := &MockExchanger{ctrl: ctrl}
	mock.recorder = &MockExchangerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchanger) EXPECT() *MockExchangerMockRecorder {
	return m.recorder
}

// GetCandlesHistory mocks base method.
func (m *MockExchanger) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCandlesHistory", pair, limit, start, end)
	ret0, _ := ret[0].(CandlesHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandlesHistory indicates an expected call of GetCandlesHistory.
func (mr *MockExchangerMockRecorder) GetCandlesHistory(pair, limit, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandlesHistory", reflect.TypeOf((*MockExchanger)(nil).GetCandlesHistory), pair, limit, start, end)
}

// GetClosePrice mocks base method.
func (m *MockExchanger) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosePrice", pair, limit, start, end)
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClosePrice indicates an expected call of GetClosePrice.
func (mr *MockExchangerMockRecorder) GetClosePrice(pair, limit, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosePrice", reflect.TypeOf((*MockExchanger)(nil).GetClosePrice), pair, limit, start, end)
}

// GetCurrencies mocks base method.
func (m *MockExchanger) GetCurrencies() (Currencies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrencies")
	ret0, _ := ret[0].(Currencies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrencies indicates an expected call of GetCurrencies.
func (mr *MockExchangerMockRecorder) GetCurrencies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrencies", reflect.TypeOf((*MockExchanger)(nil).GetCurrencies))
}

// GetOrderBook mocks base method.
func (m *MockExchanger) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{limit}
	for _, a := range pairs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrderBook", varargs...)
	ret0, _ := ret[0].(OrderBook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderBook indicates an expected call of GetOrderBook.
func (mr *MockExchangerMockRecorder) GetOrderBook(limit interface{}, pairs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{limit}, pairs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderBook", reflect.TypeOf((*MockExchanger)(nil).GetOrderBook), varargs...)
}

// GetTicker mocks base method.
func (m *MockExchanger) GetTicker() (Ticker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicker")
	ret0, _ := ret[0].(Ticker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicker indicates an expected call of GetTicker.
func (mr *MockExchangerMockRecorder) GetTicker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicker", reflect.TypeOf((*MockExchanger)(nil).GetTicker))
}

// GetTrades mocks base method.
func (m *MockExchanger) GetTrades(pairs ...string) (Trades, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range pairs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrades", varargs...)
	ret0, _ := ret[0].(Trades)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrades indicates an expected call of GetTrades.
func (mr *MockExchangerMockRecorder) GetTrades(pairs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrades", reflect.TypeOf((*MockExchanger)(nil).GetTrades), pairs...)
}

// MockIndicatorer is a mock of Indicatorer interface.
type MockIndicatorer struct {
	ctrl     *gomock.Controller
	recorder *MockIndicatorerMockRecorder
}

// MockIndicatorerMockRecorder is the mock recorder for MockIndicatorer.
type MockIndicatorerMockRecorder struct {
	mock *MockIndicatorer
}

// NewMockIndicatorer creates a new mock instance.
func NewMockIndicatorer(ctrl *gomock.Controller) *MockIndicatorer {
	mock := &MockIndicatorer{ctrl: ctrl}
	mock.recorder = &MockIndicatorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndicatorer) EXPECT() *MockIndicatorerMockRecorder {
	return m.recorder
}

// EMA mocks base method.
func (m *MockIndicatorer) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EMA", pair, resolution, period, from, to)
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EMA indicates an expected call of EMA.
func (mr *MockIndicatorerMockRecorder) EMA(pair, resolution, period, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EMA", reflect.TypeOf((*MockIndicatorer)(nil).EMA), pair, resolution, period, from, to)
}

// SMA mocks base method.
func (m *MockIndicatorer) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMA", pair, resolution, period, from, to)
	ret0, _ := ret[0].([]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SMA indicates an expected call of SMA.
func (mr *MockIndicatorerMockRecorder) SMA(pair, resolution, period, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMA", reflect.TypeOf((*MockIndicatorer)(nil).SMA), pair, resolution, period, from, to)
}