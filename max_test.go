package fp_test

import (
	"github.com/michaeljpetter/fp"
	"slices"
	"testing"
)

func TestMax(t *testing.T) {
	in := []int{6, 44, -23, 9}

	out := fp.Max(0)(slices.Values(in))

	exp := 44

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMaxInitial(t *testing.T) {
	in := []int{6, 44, -23, 9}

	out := fp.Max(50)(slices.Values(in))

	exp := 50

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMaxEmpty(t *testing.T) {
	out := fp.Max("big")(slices.Values[[]string](nil))

	exp := "big"

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMaxOf(t *testing.T) {
	in := []string{"argh", "great", "no"}

	out := fp.MaxOf(func(s string) int { return len(s) }, 3)(slices.Values(in))

	exp := 5

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMaxOfInitial(t *testing.T) {
	in := []string{"argh", "great", "no"}

	out := fp.MaxOf(func(s string) int { return len(s) }, 10)(slices.Values(in))

	exp := 10

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}

func TestMaxOfEmpty(t *testing.T) {
	out := fp.MaxOf(func(s string) int { return len(s) }, -7)(slices.Values[[]string](nil))

	exp := -7

	if out != exp {
		t.Errorf("wrong max %v, expected %v", out, exp)
	}
}
