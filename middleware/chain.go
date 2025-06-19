package middleware

import (
	"net/http"
	"slices"
)

type Chain []middleware

func (c Chain) Then(h http.Handler) http.Handler {
	for _, m := range slices.Backward(c) {
		h = m(h)
	}
	return h
}

func (c Chain) ThenFunc(h http.HandlerFunc) http.Handler {
	return c.Then(h)
}
