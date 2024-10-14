package fp

import (
	"iter"
)

// Index maps a unary sequence into a binary sequence,
// pairing each element with its zero-based index in the source.
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
