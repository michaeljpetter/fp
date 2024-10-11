package fp

import (
	"github.com/michaeljpetter/fp/types"
	"iter"
)

func Max[T types.Ordered](initial T) func(iter.Seq[T]) T {
	return Reduce(func(r T, t T) T { return max(r, t) }, initial)
}

func MaxOf[T any, U types.Ordered](project func(T) U, initial U) func(iter.Seq[T]) U {
	return Reduce(func(r U, t T) U { return max(r, project(t)) }, initial)
}
