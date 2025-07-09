package slc

import (
	"math/rand/v2"
	"reflect"
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

var max = 1 << 16

func BenchmarkToSortedAppend(b *testing.B) {
	u := make([]int, max)
	u[0] = max
	for i := 1; i < len(u); i++ {
		u[i] = rand.IntN(max)
	}
	for b.Loop() {
		s := ToSortedAppend(u)
		if s[len(s)-1] != max {
			b.Errorf("sorting wrong")
		}
	}
}

func BenchmarkToSortedClone(b *testing.B) {
	u := make([]int, max)
	u[0] = max
	for i := 1; i < len(u); i++ {
		u[i] = rand.IntN(max)
	}
	for b.Loop() {
		s := ToSortedClone(u)
		if s[len(s)-1] != max {
			b.Errorf("sorting wrong")
		}
	}
}

func BenchmarkToSortedSeq(b *testing.B) {
	u := make([]int, max)
	u[0] = max
	for i := 1; i < len(u); i++ {
		u[i] = rand.IntN(max)
	}
	for b.Loop() {
		s := slices.Sorted(slices.Values(u))
		if s[len(s)-1] != max {
			b.Errorf("sorting wrong")
		}
	}
}

func TestChunkFunk(t *testing.T) {
	f := func(a, b int) bool {
		if b == a*2 {
			return true
		}
		return false
	}

	{
		seq := ChunkFunk([]int{1, 2, 3, 6, 7, 8}, f)

		expected := [][]int{
			{1},
			{2, 3},
			{6, 7, 8},
		}
		result := slices.Collect(seq)
		if !reflect.DeepEqual(
			result,
			expected,
		) {
			t.Fatalf("expected=%v, got=%v", expected, result)
		}
	}
	{
		seq := ChunkFunk([]int{1}, f)

		expected := [][]int{
			{1},
		}
		result := slices.Collect(seq)
		if !reflect.DeepEqual(
			result,
			expected,
		) {
			t.Fatalf("expected=%v, got=%v", expected, result)
		}
	}
	{
		seq := ChunkFunk([]int{}, f)

		expected := [][]int{
			{},
		}
		result := slices.Collect(seq)
		if !reflect.DeepEqual(
			result,
			expected,
		) {
			t.Fatalf("expected=%v, got=%v", expected, result)
		}
	}
	{
		var s []int
		seq := ChunkFunk(s, f)

		var empty []int
		expected := [][]int{
			empty,
		}
		result := slices.Collect(seq)
		if !reflect.DeepEqual(
			result,
			expected,
		) {
			t.Fatalf("expected=%v, got=%v", expected, result)
		}
	}
	{
		f := func(a, b int) bool {
			return false
		}

		seq := ChunkFunk([]int{1, 2, 3, 6, 7, 8}, f)
		expected := [][]int{
			{1, 2, 3, 6, 7, 8},
		}
		result := slices.Collect(seq)
		if !reflect.DeepEqual(
			result,
			expected,
		) {
			t.Fatalf("expected=%v, got=%v", expected, result)
		}
	}
	{
		f := func(a, b int) bool {
			return true
		}

		seq := ChunkFunk([]int{1, 2, 3, 6, 7, 8}, f)
		expected := [][]int{
			{1},
			{2},
			{3},
			{6},
			{7},
			{8},
		}
		result := slices.Collect(seq)
		if !reflect.DeepEqual(
			result,
			expected,
		) {
			t.Fatalf("expected=%v, got=%v", expected, result)
		}
	}
}
