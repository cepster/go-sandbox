package main

import (
	"log"
	"net/http"
)

/*
 * Logging Middleware
 */
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hit the endpoint: %s", r.URL)
		next.ServeHTTP(w, r)
	})
}
