package fp_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestTimes(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		out := slices.Collect(fp.Times(
			func(i int) string { return strings.Repeat("X", i) },
		)(4))

		exp := []string{"", "X", "XX", "XXX"}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		out := slices.Collect(fp.Times(func(int) int { return 0 })(0))

		exp := []int{}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		out := slices.Collect(fp.Times(func(int) int { return 0 })(-3))

		exp := []int{}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})
}
