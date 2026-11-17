package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResoponseTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("Response Time Middleware")

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Response Time Middleware Returning...")
		start := time.Now()

		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		
		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String())

		fmt.Printf("Method: %s URL: %s, Status: %v, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Response Time Limiter Done")
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