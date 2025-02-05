package routes

import (
	"log"
	"net/http"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		ww := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(ww, r)
		duration := time.Since(start)
		log.Printf("Completed %s %s with status %d in %s", r.Method, r.URL.Path, ww.status, duration)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
