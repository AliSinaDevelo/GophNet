package utils

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware function that logs incoming requests.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ReponseWriter, r *http.Request)) {
		start := time.Now()
		// call the next handler
		next.ServeHTTP(w, r)
		// log
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}