package fp_test

import (
	"slices"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestReduce(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := []string{"moo", "moo", "ma", "mah"}

		out := fp.Reduce(
			func(r string, s string) string { return r + " " + s },
			"baby cow says:",
		)(slices.Values(in))

		exp := "baby cow says: moo moo ma mah"

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := fp.Reduce(
			func(string, error) string { return "irrelevant" },
			"happy",
		)(slices.Values[[]error](nil))

		exp := "happy"

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("AsCollect", func(t *testing.T) {
		in := []string{"one", "two", "three"}

		out := fp.Reduce(fp.SliceAppend, []string{})(slices.Values(in))

		if !slices.Equal(out, in) {
			t.Errorf("wrong value %v, expected %v", out, in)
		}
	})
}
