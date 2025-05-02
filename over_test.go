package fp_test

import (
	"math"
	"slices"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestOver(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := -7.2

		out := fp.Over(math.Floor, math.Trunc, math.Abs)(in)

		exp := []float64{-8, -7, 7.2}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := fp.Over[int, string]()(33)

		if 0 != len(out) {
			t.Errorf("wrong value %v, expected empty", out)
		}
	})
}

func TestOver2(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in1, in2 := 8., 3.

		out := fp.Over2(math.Min, math.Max, math.Pow)(in1, in2)

		exp := []float64{3, 8, 512}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := fp.Over2[int, string, float64]()(33, "huh?")

		if 0 != len(out) {
			t.Errorf("wrong value %v, expected empty", out)
		}
	})
}
