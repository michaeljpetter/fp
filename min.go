package fp

import (
	"github.com/michaeljpetter/fp/types"
	"iter"
)

// Min reduces a unary sequence into a scalar value,
// calculating the minimum of the given initial value and all elements in the source.
func Min[T types.Ordered](initial T) func(iter.Seq[T]) T {
	return Reduce(func(r T, t T) T { return min(r, t) }, initial)
}

// MinOf reduces a unary sequence into a scalar value,
// applying a projection to each element in the source and
// calculating the minimum of the given initial value and all projected values.
func MinOf[T any, U types.Ordered](project func(T) U, initial U) func(iter.Seq[T]) U {
	return Reduce(func(r U, t T) U { return min(r, project(t)) }, initial)
}
