package fp

import (
	"iter"
)

// Times creates a unary sequence,
// applying a projection to each index in the range [0, n).
func Times[U any](project func(int) U) func(int) iter.Seq[U] {
	return func(n int) iter.Seq[U] {
		return Map(project)(Range(0, n))
	}
}
