package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	indexHandler := http.HandlerFunc(welcome)
	mux.Handle("GET /", withLogging(indexHandler))

	log.Print("Starting simple app on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatalf("Server error: %v", err)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	t := time.Now().UTC().Format("2006-01-02T15:04:05")

	fmt.Fprintf(w, "Hello from simple app.\nCurrent time: %s", t)
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"Method: %s, Path: %s, Duration: %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
