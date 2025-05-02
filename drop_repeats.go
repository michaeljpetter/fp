package fp

import (
	"iter"
)

// DropRepeats maps a unary sequence to a unary sequence,
// dropping consecutive duplicates from the source, as determined by
// equality comparison of each element.
func DropRepeats[T comparable](src iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull(src)
		defer stop()

		t0, ok := next()
		if !ok {
			return
		}

		for yield(t0) {
			for {
				t, ok := next()
				if !ok {
					return
				}
				if t != t0 {
					t0 = t
					break
				}
			}
		}
	}
}

// DropRepeatsBy maps a unary sequence to a unary sequence,
// dropping consecutive duplicates from the source, as determined by
// equality comparison of a projected value of each element.
func DropRepeatsBy[T any, U comparable](project func(T) U) func(iter.Seq[T]) iter.Seq[T] {
	return func(src iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			next, stop := iter.Pull(src)
			defer stop()

			t0, ok := next()
			if !ok {
				return
			}

			for yield(t0) {
				u0 := project(t0)
				for {
					t, ok := next()
					if !ok {
						return
					}
					if u := project(t); u != u0 {
						t0 = t
						break
					}
				}
			}
		}
	}
}
