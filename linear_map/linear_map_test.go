package lmap

import (
	"math/rand"
	"strconv"
	"testing"
)

var size uint64 = 65536 / 2

func BenchmarkLmap(b *testing.B) {
	lm := New[int, string](size * 4)
	for b.Loop() {
		for range size {
			ii := rand.Intn(int(size))
			lm.Put(ii, strconv.Itoa(int(ii)))
		}
		for range size {
			ii := rand.Intn(int(size))
			_, _ = lm.Get(ii)
		}
		// if lm.size != int(size) {
		// 	b.Fatalf("expected %d elemns in map, got %d", size, lm.size)
		// }
		b.Log(lm.colls)
		lm.Clear()
	}
}

func BenchmarkNormalMap(b *testing.B) {
	m := make(map[int]string, size)
	for b.Loop() {
		for range size {
			ii := rand.Intn(int(size))
			m[ii] = strconv.Itoa(ii)
		}

		for range size {
			ii := rand.Intn(int(size))
			_, _ = m[ii]
		}
		// if len(m) != int(size) {
		// 	b.Fatalf("expected %d elemns in map, got %d", size, len(m))
		// }
		clear(m)
	}
}
