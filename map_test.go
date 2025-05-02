package fp_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestMap(t *testing.T) {
	in := []string{"A", "Long", "One"}

	out := slices.Collect(fp.Map(
		func(s string) int { return 2 * len(s) },
	)(slices.Values(in)))

	exp := []int{2, 8, 6}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}

func TestMap2(t *testing.T) {
	in := []string{"A", "B", "C"}

	out := slices.Collect(fp.Map2(
		func(i int, s string) string { return strings.Repeat(s, i+1) },
	)(slices.All(in)))

	exp := []string{"A", "BB", "CCC"}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}

func TestMapTo2(t *testing.T) {
	type tuple struct {
		t1 string
		t2 int
	}

	in := []string{"Ant", "Orca", "Mantis"}

	out := make([]tuple, 0)
	for out1, out2 := range fp.MapTo2(
		func(s string) (string, int) { return strings.ToUpper(s), len(s) },
	)(slices.Values(in)) {
		out = append(out, tuple{out1, out2})
	}

	exp := []tuple{
		{"ANT", 3}, {"ORCA", 4}, {"MANTIS", 6},
	}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}

func TestMap2To2(t *testing.T) {
	type tuple struct {
		t1 string
		t2 int
	}

	in := []string{"Ant", "Orca", "Mantis"}

	out := make([]tuple, 0)
	for out1, out2 := range fp.Map2To2(
		func(i int, s string) (string, int) { return s[:i+1], i * i },
	)(slices.All(in)) {
		out = append(out, tuple{out1, out2})
	}

	exp := []tuple{
		{"A", 0}, {"Or", 1}, {"Man", 4},
	}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong elements %v, expected %v", out, exp)
	}
}
