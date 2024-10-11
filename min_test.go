package fp_test

import (
	"github.com/michaeljpetter/fp"
	"slices"
	"testing"
)

func TestMin(t *testing.T) {
	in := []int{6, 44, -23, 9}

	out := fp.Min(0)(slices.Values(in))

	exp := -23

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMinInitial(t *testing.T) {
	in := []int{6, 44, -23, 9}

	out := fp.Min(-50)(slices.Values(in))

	exp := -50

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMinEmpty(t *testing.T) {
	out := fp.Min("big")(slices.Values[[]string](nil))

	exp := "big"

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMinOf(t *testing.T) {
	in := []string{"argh", "great", "no"}

	out := fp.MinOf(func(s string) int { return len(s) }, 4)(slices.Values(in))

	exp := 2

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMinOfInitial(t *testing.T) {
	in := []string{"argh", "great", "no"}

	out := fp.MinOf(func(s string) int { return len(s) }, 1)(slices.Values(in))

	exp := 1

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMinOfEmpty(t *testing.T) {
	out := fp.MinOf(func(s string) int { return len(s) }, 7)(slices.Values[[]string](nil))

	exp := 7

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}
