package lmap

import (
	"strconv"
	"testing"
)

var size uint64 = 50_000

func BenchmarkLmap(b *testing.B) {
	for b.Loop() {
		lm := New[int, string](size * 4)
		for j := range size {
			lm.Put(int(j), strconv.Itoa(int(j)))
		}
		for j := range size {
			_, _ = lm.Get(int(j))
		}
		if lm.size != int(size) {
			b.Fatalf("expected %d elemns in map, got %d", size, lm.size)
		}
	}
}

func BenchmarkNormalMap(b *testing.B) {
	for b.Loop() {
		m := make(map[int]string, size)
		for j := range size {
			m[int(j)] = strconv.Itoa(int(j))
		}

		for k := range size {
			_, _ = m[int(k)]
		}
		if len(m) != int(size) {
			b.Fatalf("expected %d elemns in map, got %d", size, len(m))
		}
	}
}
