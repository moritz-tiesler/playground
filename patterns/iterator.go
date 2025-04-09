package patterns

import "iter"

func Range(from, to int) iter.Seq[int] {

	return func(yield func(int) bool) {
		for ; from < to; from++ {
			if !yield(from) {
				return
			}
		}
	}
}

func ThreeTimes(yield func(i int) bool) {
	if !yield(1) {
		return
	}
	if !yield(2) {
		return
	}
	if !yield(3) {
		return
	}
}

func Merge[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			seq(yield)
		}
	}
}

func TakeWhile[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(v T) bool {
			if !predicate(v) {
				return false
			}
			return yield(v)
		})
	}
}

func ForEach[T any](seq iter.Seq[T], action func(T)) {
	seq(func(v T) bool {
		action(v)
		return true
	})
}
