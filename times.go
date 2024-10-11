package fp

import (
	"iter"
)

func Times[U any](project func(int) U) func(int) iter.Seq[U] {
	return func(n int) iter.Seq[U] {
		return Map(project)(Range(0, n))
	}
}
