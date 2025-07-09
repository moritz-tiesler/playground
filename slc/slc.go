package slc

import (
	"cmp"
	"iter"
	"slices"
)

func ToSortedAppend[T cmp.Ordered](ss []T) []T {
	var o []T
	if ss == nil {
		return nil
	}
	o = append(o, ss...)
	slices.Sort(o)
	return o
}

func ToSortedClone[T cmp.Ordered](ss []T) []T {
	var o []T
	if ss == nil {
		return nil
	}
	o = make([]T, len(ss))
	copy(o, ss)
	slices.Sort(o)
	return o
}

// ChunkFunc returns an iterator over consecutive sub-slices.
// When the funcion passed to ChunkFunc returns true for a given a and b,
// a cut will be made be made after the index of a and a chunk will be returned.
func ChunkFunk[T any](s []T, doCut func(a, b T) bool) iter.Seq[[]T] {
	chunk := []T{}
	left, right := 0, 1
	return func(yield func([]T) bool) {
		if len(s) < 2 {
			yield(s)
			return
		}
		for {
			if left >= len(s) {
				return
			}
			if right >= len(s) {
				end := min(right, len(s))
				chunk = s[left:end:end]
				yield(chunk)
				return
			}
			if !doCut(s[right-1], s[right]) {
				right++
				continue
			} else {
				chunk = s[left:right:right]
				if !yield(chunk) {
					return
				}
				left = right
				right++
			}
		}
	}
}
