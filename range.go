package fp

import (
	"github.com/michaeljpetter/fp/types"
	"iter"
)

func Range[T types.Number](start, end T) iter.Seq[T] {
	return RangeStep(start, end, 1)
}

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
