package lmap

import (
	"strconv"
	"testing"
)

func BenchmarkLmap(b *testing.B) {
	var size uint64 = 10_000
	for b.Loop() {
		lm := New[int, string](size * 4)
		for j := range size {
			lm.Put(int(j), strconv.Itoa(int(j)))
		}
		for j := range size {
			_, _ = lm.Get(int(j))
		}
		if lm.size != 10_000 {
			b.Fatalf("expected %d elemns in map, got %d", size, lm.size)
		}
		b.Logf("performed hops=%d", lm.hops)
		b.Logf("num hashes generated=%d", len(lm.hashesUsed))
		b.Logf("num trunc generated=%d", len(lm.truncHashesUsed))
	}
}

func BenchmarkNormalMap(b *testing.B) {
	var size uint64 = 10_000
	for b.Loop() {
		m := make(map[int]string, size)
		for j := range size {
			m[int(j)] = strconv.Itoa(int(j))
		}

		for k := range size {
			_, _ = m[int(k)]
		}
		if len(m) != 10_000 {
			b.Fatalf("expected %d elemns in map, got %d", size, len(m))
		}
	}
}
