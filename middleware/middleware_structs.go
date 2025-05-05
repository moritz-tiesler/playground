package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type LoggingMiddleware struct {
	h http.Handler
}

func (l LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v: LOG %s - %s %s %s\n",
		time.Now(), r.RemoteAddr, r.Proto, r.Method, r.URL)
	l.h.ServeHTTP(w, r)
}

func NewLoggingMiddleware(h http.Handler) *LoggingMiddleware {
	return &LoggingMiddleware{h}
}
