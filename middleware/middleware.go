package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func NewServer() *http.Server {
	router := http.NewServeMux()
	base := Chain{LogRequestMiddleware(fmt.Printf)}
	router.Handle("/", NewLoggingMiddleware(http.HandlerFunc(home)))
	router.Handle("/bye", base.ThenFunc(bye))
	stack := Stack(
		SecureHeadersMiddleware(map[string]string{
			"X":               "1; mode-block",
			"X-Frame-Options": "deny",
		}),
	)
	server := &http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	return server
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v: Welcome to Go Middleware\n", time.Now())))
}

func bye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye from Go Middleware\n"))
}
func LogRequestMiddleware(loggingFunc func(string, ...any) (int, error)) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loggingFunc("%v: LOG %s - %s %s %s\n", time.Now(), r.RemoteAddr, r.Proto, r.Method, r.URL)

			next.ServeHTTP(w, r)
		})
	}
}

func SecureHeadersMiddleware(headers map[string]string) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for k, v := range headers {
				// w.Header().Set("X-XSS-Protection", "1; mode-block")
				w.Header().Set(k, v)
			}

			next.ServeHTTP(w, r)
		})
	}
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
