package fp

import (
	"iter"
)

// Reduce reduces a unary sequence into a scalar value,
// sequentially pairing an accumulator value with each element in the source,
// applying the given reducer function to each pair, and using its result as
// the accumulator value for the next pair.
// The given initial value is used as the first accumulator value, and
// the final accumulator value returned by the reducer is returned by Reduce.
func Reduce[T, R any](reducer func(R, T) R, initial R) func(iter.Seq[T]) R {
	return func(src iter.Seq[T]) R {
		r := initial
		for t := range src {
			r = reducer(r, t)
		}
		return r
	}
}
