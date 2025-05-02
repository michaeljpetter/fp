package fp

import (
	"iter"
)

// FLatten iterates a unary sequence of nested unary sub-sequences as a single unnested unary sequence.
// The iteration order is stable and depth-first.
func Flatten[T any](src iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for seq := range src {
			for t := range seq {
				if !yield(t) {
					return
				}
			}
		}
	}
}

// FLatMap maps a unary sequence into a unary sequence,
// applying a sequence projection to each element in the source and
// then flattening the resulting sequences as with [Flatten].
// The iteration order is stable and depth-first.
func FlatMap[T, U any](project func(T) iter.Seq[U]) func(iter.Seq[T]) iter.Seq[U] {
	return func(src iter.Seq[T]) iter.Seq[U] {
		return func(yield func(U) bool) {
			for t := range src {
				for u := range project(t) {
					if !yield(u) {
						return
					}
				}
			}
		}
	}
}
