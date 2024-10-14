package fp

import (
	"github.com/michaeljpetter/fp/types"
	"iter"
)

// Range creates a unary sequence of range [start, end) with a step of 1.
// If end is less than start, the result will be empty.
func Range[T types.Number](start, end T) iter.Seq[T] {
	return RangeStep(start, end, 1)
}

// Range creates a unary sequence of range [start, end) with the given step size.
// If step is positive, the sequence will increase from start to end.
// If step is negative, the sequence will decrease from start to end.
// If step is positive and end is less than start, the result will be empty.
// If step is negative and start is less than end, the result will be empty.
// If step is zero, Range panics with [ErrZero].
func RangeStep[T types.Number](start, end, step T) iter.Seq[T] {
	if 0 < step {
		return func(yield func(T) bool) {
			for n := start; n < end; n += step {
				if !yield(n) {
					return
				}
			}
		}
	}

	if step < 0 {
		return func(yield func(T) bool) {
			for n := start; end < n; n += step {
				if !yield(n) {
					return
				}
			}
		}
	}

	panic(newErrZero("step"))
}
