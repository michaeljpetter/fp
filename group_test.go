package fp_test

import (
	"iter"
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestGroupBy(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		out := maps.Collect(fp.Map2To2(fp.Project2To2(fp.Identity[int], slices.Collect[int]))(
			fp.GroupBy(
				func(i int) int { return i % 3 },
			)(slices.Values(in))))

		exp := map[int][]int{
			0: {3, 6, 9},
			1: {1, 4, 7, 10},
			2: {2, 5, 8},
		}

		if !maps.EqualFunc(out, exp, slices.Equal) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := maps.Collect(fp.Map2To2(fp.Project2To2(fp.Identity[int], slices.Collect[int]))(
			fp.GroupBy(
				func(i int) int { return i % 3 },
			)(slices.Values([]int(nil)))))

		if 0 != len(out) {
			t.Errorf("wrong elements %v, expected empty", out)
		}
	})
}

func TestGroupStableBy(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		out := maps.Collect(fp.Map2To2(fp.Project2To2(fp.Identity[int], slices.Collect[int]))(
			fp.GroupStableBy(
				func(i int) int { return i % 3 },
			)(slices.Values(in))))

		exp := map[int][]int{
			0: {3, 6, 9},
			1: {1, 4, 7, 10},
			2: {2, 5, 8},
		}

		if !maps.EqualFunc(out, exp, slices.Equal) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Stable", func(t *testing.T) {
		in := strings.Split("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz", "")

		out := slices.Collect(fp.Map2(func(s string, _ iter.Seq[string]) string { return s })(
			fp.GroupStableBy(
				strings.ToUpper,
			)(slices.Values(in))))

		exp := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

		if !slices.Equal(out, exp) {
			t.Errorf("wrong elements %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := maps.Collect(fp.Map2To2(fp.Project2To2(fp.Identity[int], slices.Collect[int]))(
			fp.GroupStableBy(
				func(i int) int { return i % 3 },
			)(slices.Values([]int(nil)))))

		if 0 != len(out) {
			t.Errorf("wrong elements %v, expected empty", out)
		}
	})
}
