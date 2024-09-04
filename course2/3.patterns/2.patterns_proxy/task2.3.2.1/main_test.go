package main

import (
	"reflect"
	"testing"
)

type MockLines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (t *MockLines) StochPrice() ([]float64, []float64) {
	a := []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456}
	b := []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546}
	return a, b
}
func (t *MockLines) RSI(period int) ([]float64, []float64) {
	a := []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456}
	b := []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546}
	return a, b
}
func (t *MockLines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	a := []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456}
	b := []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546}
	return a, b
}

func (t *MockLines) MACD() ([]float64, []float64) {
	a := []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456}
	b := []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546}
	return a, b
}
func (t *MockLines) EMA() []float64 {
	a := []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456}
	return a
}
func (t *MockLines) SMA(period int) []float64 {
	a := []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456}
	return a
}

func MockLoadKlinesProxy() *LinesProxy {
	return &LinesProxy{
		lines: &MockLines{},
		cache: make(map[string][]float64),
	}
}

func TestLinesProxy_StochPrice(t *testing.T) {

	tests := []struct {
		name  string
		want  []float64
		want1 []float64
	}{
		{name: "StochPrice",
			want:  []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456},
			want1: []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lp := MockLoadKlinesProxy()

			got, got1 := lp.StochPrice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StochPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StochPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_RSI(t *testing.T) {

	tests := []struct {
		name  string
		want  []float64
		want1 []float64
	}{
		{name: "RSI",
			want:  []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456},
			want1: []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lp := MockLoadKlinesProxy()

			got, got1 := lp.RSI(4)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StochPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StochPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_StochRSI(t *testing.T) {

	tests := []struct {
		name  string
		want  []float64
		want1 []float64
	}{
		{name: "StochRSI",
			want:  []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456},
			want1: []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lp := MockLoadKlinesProxy()

			got, got1 := lp.StochRSI(4)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StochPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StochPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_MACD(t *testing.T) {

	tests := []struct {
		name  string
		want  []float64
		want1 []float64
	}{
		{name: "MACD",
			want:  []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456},
			want1: []float64{0.135453, 0.165464, 0.13459, 1.46734554, 2.56546},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lp := MockLoadKlinesProxy()

			got, got1 := lp.MACD()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StochPrice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("StochPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLinesProxy_EMA(t *testing.T) {

	tests := []struct {
		name string
		want []float64
	}{
		{name: "EMA",
			want: []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lp := MockLoadKlinesProxy()

			got := lp.EMA()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StochPrice() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestLinesProxy_SMA(t *testing.T) {

	tests := []struct {
		name string
		want []float64
	}{
		{name: "SMA",
			want: []float64{0.15765, 0.245353, 0.123544, 1.1456, 2.5263456},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lp := MockLoadKlinesProxy()

			got := lp.SMA(1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StochPrice() got = %v, want %v", got, tt.want)
			}

		})
	}
}
