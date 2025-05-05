package middleware

import (
	"fmt"
	"net/http"
)

func NewServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	stack := Stack(LogRequestMiddleware, SecureHeadersMiddleware)
	server := &http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	return server
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Middleware\n"))
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

type middleware func(http.Handler) http.Handler

func Stack(middlewares ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i > -1; i-- {
			m := middlewares[i]
			next = m(next)
		}
		return next
	}
}

// Stack(log, secure, home)

// log(secure(home))
