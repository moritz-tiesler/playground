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

func Cycle[T any](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func TakeWhile[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if !predicate(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

func ForEach[T any](seq iter.Seq[T], action func(T)) {
	for v := range seq {
		action(v)
	}
}

func PrimeSieve(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		sieve := make([]bool, limit+1)
		for i := 2; i <= limit; i++ {
			if !sieve[i] {
				if !yield(i) {
					return
				}
				for j := i * i; j <= limit; j += i {
					sieve[j] = true
				}
			}
		}
	}
}
