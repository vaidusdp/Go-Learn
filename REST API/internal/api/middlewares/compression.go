package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	fmt.Println("Compression Middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compression Middleware Returning...")

		if !strings.Contains(r.Header.Get("accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}

		next.ServeHTTP(w, r)
		fmt.Println("Sent response from compresion middleware")
		fmt.Println("Compression Middleware Done")
	})
}
 
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}