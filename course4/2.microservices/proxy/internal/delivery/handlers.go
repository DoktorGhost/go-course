package delivery

import (
	"io"
	"net/http"
	"net/url"
	"proxy/internal/broker"
	"proxy/internal/limiter"
)

func ProxyHandler(target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxyURL, err := url.Parse(target)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}

		r.URL.Scheme = proxyURL.Scheme
		r.URL.Host = proxyURL.Host

		resp, err := http.DefaultTransport.RoundTrip(r)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientIP string
		if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
			clientIP = realIP
		} else if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			clientIP = forwardedFor
		} else {
			clientIP = r.RemoteAddr
		}

		limiter := limiter.GetLimiterForIP(clientIP)

		if !limiter.Limiter.Take() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("429 Too Many Requests - rate limit exceeded"))
			broker.SendMessage("Rate limit exceeded for IP: " + clientIP)
			return
		}

		next(w, r)
	}
}
