package slc

import (
	"cmp"
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
