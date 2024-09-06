package main

const (
	ProductCocaCola = iota
	ProductPepsi
	ProductSprite
)

type Product struct {
	ProductID     int
	Sells         []float64
	Buys          []float64
	CurrentPrice  float64
	ProfitPercent float64
}

type Profitable interface {
	SetProduct(p *Product)
	GetAverageProfit() float64
	GetAverageProfitPercent() float64
	GetCurrentProfit() float64
	GetDifferenceProfit() float64
	GetAllData() []float64
	Average(price []float64) float64
	Sum(price []float64) float64
}

type StatisticProfit struct {
	product                 *Product
	getAverageProfit        func() float64
	getAverageProfitPercent func() float64
	getCurrentProfit        func() float64
	getDifferenceProfit     func() float64
	getAllData              func() []float64
}

func NewStatisticProfit(opts ...func(profit *StatisticProfit)) Profitable {
	profit := &StatisticProfit{}

	for _, opt := range opts {
		opt(profit)
	}

	return profit
}

// WithAverageProfit sets the average profit of the product
// AverageProfit = Average Sells - Average Buys
func WithAverageProfit(s *StatisticProfit) {
	s.getAverageProfit = func() float64 {
		if s.product == nil {
			return 0.0
		}
		return s.Average(s.product.Sells) - s.Average(s.product.Buys)
	}
}

// WithAverageProfitPercent sets the average profit percent of the product
// Average Profit Percent = Average Profit / Average Buys * 100
func WithAverageProfitPercent(s *StatisticProfit) {
	s.getAverageProfitPercent = func() float64 {
		if s.product == nil {
			return 0.0
		}
		return (s.GetAverageProfit() / s.Average(s.product.Buys)) * 100
	}
}

// WithCurrentProfit sets the current profit of the product
// Current Profit = Current Price - Current Price * (100 - Profit Percent) / 100
func WithCurrentProfit(s *StatisticProfit) {
	s.getCurrentProfit = func() float64 {
		return s.product.CurrentPrice - (s.product.CurrentPrice * (100 - s.product.ProfitPercent) / 100)
	}
}

// WithDifferenceProfit sets the difference profit of the product
// Difference Profit = Current Price - Average Sells
func WithDifferenceProfit(s *StatisticProfit) {
	s.getDifferenceProfit = func() float64 {
		avgSells := s.Average(s.product.Sells)
		return s.product.CurrentPrice - avgSells
	}
}

func WithAllData(s *StatisticProfit) {
	s.getAllData = func() []float64 {
		res := make([]float64, 0, 4)
		if s.getAverageProfitPercent != nil {
			res = append(res, s.getAverageProfitPercent())
		}
		if s.getCurrentProfit != nil {
			res = append(res, s.getCurrentProfit())
		}
		if s.getDifferenceProfit != nil {
			res = append(res, s.getDifferenceProfit())
		}
		return res
	}
}

// реализация интерфейса
// устанавливаем продукт для вычислений
func (s *StatisticProfit) SetProduct(p *Product) {
	s.product = p
}

// вычисляем сумму значений для среза значений
func (s *StatisticProfit) Sum(price []float64) float64 {
	if len(price) == 0 {
		return 0
	}
	var sum float64
	for _, v := range price {
		sum += v
	}
	return sum
}

// вычисляем среднее значение
func (s *StatisticProfit) Average(price []float64) float64 {
	if len(price) == 0 {
		return 0
	}
	return s.Sum(price) / float64(len(price))
}

// возвращаем среднюю прибыль
func (s *StatisticProfit) GetAverageProfit() float64 {
	return s.getAverageProfit()
}

// возвращаем средний процент прибыли
func (s *StatisticProfit) GetAverageProfitPercent() float64 {
	return s.getAverageProfitPercent()
}

// возвращаем текущую прибыль
func (s *StatisticProfit) GetCurrentProfit() float64 {
	return s.getCurrentProfit()
}

// возвращаем прибыль от разницы
func (s *StatisticProfit) GetDifferenceProfit() float64 {
	return s.getDifferenceProfit()
}

// возвращаем метод структуры
func (s *StatisticProfit) GetAllData() []float64 {
	return s.getAllData()

}

func main() {
	product := &Product{
		ProductID:     ProductCocaCola,
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}

	product.ProductID = ProductCocaCola
	product.ProductID = ProductPepsi
	product.ProductID = ProductSprite

	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)

	statProfit.SetProduct(product)
}
