package stringbuild

import (
	"math/rand"
	"testing"
)

var listSize = 10000
var strSize = 20

func BenchmarkWithPlus(b *testing.B) {
	testData := makeTestData(listSize, strSize)
	for i := 0; i < b.N; i++ {
		RunConcat(testData, WithPlus)
	}
}

func BenchmarkWithBuilder(b *testing.B) {
	testData := makeTestData(listSize, strSize)
	for i := 0; i < b.N; i++ {
		RunConcat(testData, WithBuilder)
	}
}

func makeTestData(listSize int, strSize int) []string {
	testData := make([]string, listSize)
	for i := 0; i < listSize; i++ {
		testData[i] = RandStringRunes(strSize)
	}
	return testData
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringRunes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
