package fp_test

import (
	"slices"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestDropRepeats(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := []string{"moo", "moo", "nah", "nah", "coco"}

		out := slices.Collect(fp.DropRepeats(slices.Values(in)))

		exp := []string{"moo", "nah", "coco"}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("NoRepeats", func(t *testing.T) {
		in := []int{1, 6, 1, 12}

		out := slices.Collect(fp.DropRepeats(slices.Values(in)))

		if !slices.Equal(out, in) {
			t.Errorf("wrong elements %v, expected %v", out, in)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := slices.Collect(fp.DropRepeats(slices.Values([]int(nil))))

		if 0 != len(out) {
			t.Errorf("wrong elements %v, expected empty", out)
		}
	})

	t.Run("AllRepeats", func(t *testing.T) {
		out := slices.Collect(fp.DropRepeats(slices.Values([]int(nil))))

		if 0 != len(out) {
			t.Errorf("wrong elements %v, expected empty", out)
		}
	})
}

func TestDropRepeatsBy(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := []string{"moo", "ma", "no", "nah", "coco"}

		out := slices.Collect(fp.DropRepeatsBy(
			func(s string) byte { return s[0] },
		)(slices.Values(in)))

		exp := []string{"moo", "no", "coco"}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("NoRepeats", func(t *testing.T) {
		in := []int{12, 12, 12, 12}

		out := slices.Collect(fp.DropRepeatsBy(
			func() func(int) int {
				i := 77
				return func(int) int { i++; return i }
			}(),
		)(slices.Values(in)))

		if !slices.Equal(out, in) {
			t.Errorf("wrong elements %v, expected %v", out, in)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := slices.Collect(fp.DropRepeatsBy(
			func(i int) int { return i % 2 },
		)(slices.Values([]int(nil))))

		if 0 != len(out) {
			t.Errorf("wrong elements %v, expected empty", out)
		}
	})

	t.Run("AllRepeats", func(t *testing.T) {
		in := []int{43, 3, 11, 21, 77}

		out := slices.Collect(fp.DropRepeatsBy(
			func(i int) int { return i % 2 },
		)(slices.Values(in)))

		exp := []int{43}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})
}
