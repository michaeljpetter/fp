package fp

import (
	"github.com/michaeljpetter/fp/types"
	"iter"
)

// Max reduces a unary sequence into a scalar value,
// calculating the maximum of the given initial value and all elements in the source.
func Max[T types.Ordered](initial T) func(iter.Seq[T]) T {
	return Reduce(func(r T, t T) T { return max(r, t) }, initial)
}

// MaxOf reduces a unary sequence into a scalar value,
// applying a projection to each element in the source and
// calculating the maximum of the given initial value and all projected values.
func MaxOf[T any, U types.Ordered](project func(T) U, initial U) func(iter.Seq[T]) U {
	return Reduce(func(r U, t T) U { return max(r, project(t)) }, initial)
}
