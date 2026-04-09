package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!")
}

func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Pass request to the next handler
			next.ServeHTTP(w, r)

			// Log request details
			logger.Info("http_request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remote_addr", r.RemoteAddr),
				zap.Duration("duration", time.Since(start)),
				zap.String("user_agent", r.UserAgent()))
		})
	}
}

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	loggedRouter := LoggingMiddleware(logger)(http.DefaultServeMux)
	http.ListenAndServe(":"+port, loggedRouter)
}
