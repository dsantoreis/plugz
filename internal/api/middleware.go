package api

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func JSONLoggingMiddleware(next http.Handler) http.Handler {
	logger := log.New(os.Stdout, "", 0)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		payload := map[string]any{
			"ts":          time.Now().UTC().Format(time.RFC3339Nano),
			"method":      r.Method,
			"path":        r.URL.Path,
			"status":      sw.status,
			"latency_ms":  time.Since(start).Milliseconds(),
			"remote_addr": r.RemoteAddr,
		}
		blob, _ := json.Marshal(payload)
		logger.Println(string(blob))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	expected := os.Getenv("SKILLS_API_TOKEN")
	if expected == "" {
		expected = "dev-token"
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}
		if r.Header.Get("Authorization") != "Bearer "+expected {
			respondErr(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RateLimitMiddleware(rps int, window time.Duration) func(http.Handler) http.Handler {
	_ = window
	var (
		mu       sync.Mutex
		limiters = map[string]*rate.Limiter{}
	)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			if ip == "" {
				ip = r.RemoteAddr
			}

			mu.Lock()
			limiter, ok := limiters[ip]
			if !ok {
				limiter = rate.NewLimiter(rate.Every(time.Second/time.Duration(rps)), rps)
				limiters[ip] = limiter
			}
			mu.Unlock()

			if !limiter.Allow() {
				respondErr(w, http.StatusTooManyRequests, "rate limit exceeded")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
