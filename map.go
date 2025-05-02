package fp

import (
	"iter"
)

// Map maps a unary sequence into a unary sequence,
// applying a projection to each element the source.
func Map[T, U any](project func(T) U) func(iter.Seq[T]) iter.Seq[U] {
	return func(src iter.Seq[T]) iter.Seq[U] {
		return func(yield func(U) bool) {
			for t := range src {
				if !yield(project(t)) {
					return
				}
			}
		}
	}
}

// Map2 maps a binary sequence into a unary sequence,
// applying a 2-to-1 projection to each element in the source.
func Map2[T1, T2, U any](project func(T1, T2) U) func(iter.Seq2[T1, T2]) iter.Seq[U] {
	return func(src iter.Seq2[T1, T2]) iter.Seq[U] {
		return func(yield func(U) bool) {
			for t1, t2 := range src {
				if !yield(project(t1, t2)) {
					return
				}
			}
		}
	}
}

// MapTo2 maps a unary sequence into a binary sequence,
// applying a 1-to-2 projection to each element in the source.
func MapTo2[T, U1, U2 any](project func(T) (U1, U2)) func(iter.Seq[T]) iter.Seq2[U1, U2] {
	return func(src iter.Seq[T]) iter.Seq2[U1, U2] {
		return func(yield func(U1, U2) bool) {
			for t := range src {
				if !yield(project(t)) {
					return
				}
			}
		}
	}
}

// Map2To2 maps a binary sequence into a binary sequence,
// applying a 2-to-2 projection to each element in the source.
func Map2To2[T1, T2, U1, U2 any](project func(T1, T2) (U1, U2)) func(iter.Seq2[T1, T2]) iter.Seq2[U1, U2] {
	return func(src iter.Seq2[T1, T2]) iter.Seq2[U1, U2] {
		return func(yield func(U1, U2) bool) {
			for t1, t2 := range src {
				if !yield(project(t1, t2)) {
					return
				}
			}
		}
	}
}
