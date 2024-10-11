package fp

import (
	"iter"
)

func Reduce[T, R any](reducer func(R, T) R, initial R) func(iter.Seq[T]) R {
	return func(src iter.Seq[T]) R {
		r := initial
		for t := range src {
			r = reducer(r, t)
		}
		return r
	}
}
