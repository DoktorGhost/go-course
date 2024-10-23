package metrics

import (
	"net/http"
	"time"
)

// MetricsMiddleware собирает метрики для HTTP запросов
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Обертываем ResponseWriter для сбора статуса
		wrappedWriter := &responseWriterWrapper{ResponseWriter: w}
		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(start).Seconds()
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		RequestCount.WithLabelValues(r.Method, r.URL.Path).Inc()
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
