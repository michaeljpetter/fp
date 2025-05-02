package fp_test

import (
	"slices"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestIndex(t *testing.T) {
	type tuple struct {
		t1 int
		t2 string
	}

	in := []string{"A", "Z", "O"}

	out := make([]tuple, 0)
	for out1, out2 := range fp.Index(slices.Values(in)) {
		out = append(out, tuple{out1, out2})
	}

	exp := []tuple{
		{0, "A"}, {1, "Z"}, {2, "O"},
	}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}

func TestIndexMap(t *testing.T) {
	in := []float64{3.14, 2.18, -1}

	out := slices.Collect(fp.IndexMap(
		func(i int, f float64) int { return i * int(f) },
	)(slices.Values(in)))

	exp := []int{0, 2, -2}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}
