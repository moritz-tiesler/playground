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
