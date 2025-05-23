package fp_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestRange(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		out := slices.Collect(fp.Range(-2, 3))

		exp := []int{-2, -1, 0, 1, 2}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("EndBeforeStart", func(t *testing.T) {
		out := slices.Collect(fp.Range(3, -2))

		exp := []int{}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})
}

func TestRangeStep(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		out := slices.Collect(fp.RangeStep(-2.1, 4.2, 2.1))

		exp := []float64{-2.1, 0, 2.1}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		out := slices.Collect(fp.RangeStep(4, -2, -2))

		exp := []int{4, 2, 0}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("EndBeforeStart", func(t *testing.T) {
		out := slices.Collect(fp.RangeStep(3, -2, 2))

		exp := []int{}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("EndBeforeStartNegative", func(t *testing.T) {
		out := slices.Collect(fp.RangeStep(-2, 5, -1))

		exp := []int{}

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		defer func() {
			p := recover()
			err, _ := p.(error)
			if err == nil || !errors.Is(err, fp.ErrZero) {
				t.Errorf("wrong panic %v, expected %v", p, fp.ErrZero)
			}
		}()

		fp.RangeStep(4, -2, 0)
	})
}
