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
