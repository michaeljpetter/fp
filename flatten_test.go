package fp_test

import (
	"iter"
	"slices"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestFlatten(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := [][]int{{1, 2}, {3, 4, 5}, {6}, {7, 8, 9, 10}}

		out := slices.Collect(fp.Flatten(fp.Map(slices.Values[[]int])(slices.Values(in))))

		exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := slices.Collect(fp.Flatten(slices.Values([]iter.Seq[string](nil))))

		if 0 != len(out) {
			t.Errorf("wrong elements %v, expected empty", out)
		}
	})
}

func TestFlatMap(t *testing.T) {
	in := []string{"Bob", "Jim", "Steve"}

	out := slices.Collect(fp.FlatMap(
		func(s string) iter.Seq[string] {
			return slices.Values([]string{strings.ToLower(s), strings.ToUpper(s)})
		},
	)(slices.Values(in)))

	exp := []string{"bob", "BOB", "jim", "JIM", "steve", "STEVE"}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}
