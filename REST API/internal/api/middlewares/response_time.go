package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResoponseTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		start := time.Now()

		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		
		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String())

		fmt.Printf("Method: %s URL: %s, Status: %v, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response Writer")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw * responseWriter) WriteHeader(code int){
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}