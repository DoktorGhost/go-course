package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Время запроса каждого эндпоинта
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of latencies for HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Количество запросов каждого эндпоинта
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint"},
	)

	// Время обращения в кэш
	CacheDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cache_duration_seconds",
			Help:    "Histogram of latencies for cache calls.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	// Время обращения в БД
	DbDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_duration_seconds",
			Help:    "Histogram of latencies for DB calls.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	// Время обращения во внешний API
	ApiDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_duration_seconds",
			Help:    "Histogram of latencies for external API calls.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func Init() {
	// Регистрация метрик
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(CacheDuration)
	prometheus.MustRegister(DbDuration)
	prometheus.MustRegister(ApiDuration)
}

// Handler для Prometheus
func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
