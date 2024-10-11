package fp_test

import (
	"github.com/michaeljpetter/fp"
	"slices"
	"testing"
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
