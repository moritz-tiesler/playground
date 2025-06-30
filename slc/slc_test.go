package slc

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestToSortedNil(t *testing.T) {
	var s []int

	sorted := ToSortedAppend(s)

	if sorted != nil {
		t.Errorf("expected sorted nil slice to be nil\n")
	}
	sorted = ToSortedClone(s)

	if sorted != nil {
		t.Errorf("expected sorted nil slice to be nil\n")
	}
}

func TestToSorted(t *testing.T) {
	s := []int{3, 2, 1}

	sorted := ToSortedAppend(s)
	if len(sorted) != len(s) {
		t.Errorf("ToSorted changed len")
	}

	sorted = ToSortedClone(s)
	if len(sorted) != len(s) {
		t.Errorf("ToSorted changed len, expected=%d, got=%d", len(s), len(sorted))
	}
	s[0] = -1
	if slices.Contains(sorted, -1) {
		t.Error("expected sorted slice not to be modified")
	}
}

func BenchmarkToSortedAppend(b *testing.B) {
	max := 1 << 15
	u := make([]int, max)
	for i := 0; i < len(u); i++ {
		u[i] = rand.IntN(max)
	}
	for b.Loop() {
		_ = ToSortedAppend(u)
	}
}

func BenchmarkToSortedClone(b *testing.B) {
	max := 1 << 15
	u := make([]int, max)
	for i := 0; i < len(u); i++ {
		u[i] = rand.IntN(max)
	}
	for b.Loop() {
		_ = ToSortedClone(u)
	}
}
