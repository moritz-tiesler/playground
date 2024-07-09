package generics

import "errors"

type nonEmptySlice[T any] []T

func NewNonEmptySlice[T any](s []T) (nonEmptySlice[T], error) {
	if len(s) < 1 {
		return nonEmptySlice[T]{}, errors.New("could not create nonEmptySlice; slice is empty")
	}
	l := len(s)
	nonEmpty := make(nonEmptySlice[T], l)
	for i := 0; i < l; i++ {
		nonEmpty[i] = s[i]
	}

	return nonEmpty, nil
}
