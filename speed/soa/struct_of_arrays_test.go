package soa

import (
	"testing"
)

var size = 100000

func BenchmarkArrayOfStructs(b *testing.B) {
	testData := NewArrayOfStructs()
	for i := 0; i < b.N; i++ {
		LoopArrayRandom(testData)
	}
}

func BenchmarkStructOfArrays(b *testing.B) {
	testData := NewStructOfArrays()
	for i := 0; i < b.N; i++ {
		LoopStructRandom(testData)
	}
}
