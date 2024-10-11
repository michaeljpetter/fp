package fp_test

import (
	"github.com/michaeljpetter/fp"
	"slices"
	"strings"
	"testing"
)

func TestTimes(t *testing.T) {
	out := slices.Collect(fp.Times(
		func(i int) string { return strings.Repeat("X", i) },
	)(4))

	exp := []string{"", "X", "XX", "XXX"}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}

func TestTimesZero(t *testing.T) {
	out := slices.Collect(fp.Times(func(int) int { return 0 })(0))

	exp := []int{}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}

func TestTimesNegative(t *testing.T) {
	out := slices.Collect(fp.Times(func(int) int { return 0 })(-3))

	exp := []int{}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}
