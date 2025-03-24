package middleware

import (
	"fmt"
	"net/http"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", home)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Middleware"))
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

func SecureHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode-block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}
