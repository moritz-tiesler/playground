package patterns

import (
	"reflect"
	"sync"
)

var cache map[any]any = make(map[any]any)

var cacheMux sync.Mutex

func NewSingleTon[T any]() *T {
	var zero [0]T
	hash := reflect.TypeOf(zero).Elem()

	cacheMux.Lock()
	defer cacheMux.Unlock()
	v, ok := cache[hash]

	if ok {
		return v.(*T)
	}

	v = new(T)
	cache[hash] = v
	return v.(*T)
}
