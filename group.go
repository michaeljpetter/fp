package fp

import (
	"iter"
	"slices"
)

// GroupBy transforms a unary sequence into a binary sequence of value/sub-sequence pairs,
// by applying a projection to each element in the source, then
// using the projected values to partition the elements into sub-sequences.
// The iteration order of value/sub-sequence pairs is not specified,
// though the iteration order of individual sub-sequences is guaranteed to be stable
// (will always match the order of elements in the source).
// If full stability is required, use [GroupStableBy].
func GroupBy[T any, U comparable](project func(T) U) func(iter.Seq[T]) iter.Seq2[U, iter.Seq[T]] {
	return func(src iter.Seq[T]) iter.Seq2[U, iter.Seq[T]] {
		return func(yield func(U, iter.Seq[T]) bool) {
			groups := make(map[U][]T)

			for t := range src {
				u := project(t)
				groups[u] = append(groups[u], t)
			}

			for u, g := range groups {
				if !yield(u, slices.Values(g)) {
					return
				}
			}
		}
	}
}

// GroupStableBy transforms a unary sequence into a binary sequence of value/sub-sequence pairs,
// by applying a projection to each element in the source, then
// using the projected values to partition the elements into sub-sequences.
// The iteration order of both value/sub-sequence pairs and individual sub-sequences is guaranteed to be stable
// (will always match the order of elements in the source).
func GroupStableBy[T any, U comparable](project func(T) U) func(iter.Seq[T]) iter.Seq2[U, iter.Seq[T]] {
	return func(src iter.Seq[T]) iter.Seq2[U, iter.Seq[T]] {
		return func(yield func(U, iter.Seq[T]) bool) {
			groups := make(map[U][]T)
			order := make([]U, 0)

			for t := range src {
				u := project(t)
				g, ok := groups[u]
				groups[u] = append(g, t)
				if !ok {
					order = append(order, u)
				}
			}

			for _, u := range order {
				if !yield(u, slices.Values(groups[u])) {
					return
				}
			}
		}
	}
}
