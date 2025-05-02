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

// IndexMap maps a unary sequence into a unary sequence,
// pairing each element with its zero-based index in the source and
// applying a 2-to-1 projection to each pair.
func IndexMap[T, U any](project func(int, T) U) func(iter.Seq[T]) iter.Seq[U] {
	return func(src iter.Seq[T]) iter.Seq[U] {
		return func(yield func(U) bool) {
			i := 0
			for t := range src {
				if !yield(project(i, t)) {
					return
				}
				i++
			}
		}
	}
}
