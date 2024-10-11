package fp

import (
	"iter"
)

func Index[T any](src iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for t := range src {
			if !yield(i, t) {
				return
			}
			i++
		}
	}
}
