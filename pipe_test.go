package fp_test

import (
	"slices"
	"strings"
	"testing"

	. "github.com/michaeljpetter/fp"
)

func TestPipe(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := "pipe   like  a       pro"

		out := Pipe[func(string) string](
			strings.Fields,
			slices.All[[]string],
			Map2(Pipe[func(int, string) string](
				Project2To2(
					Pipe[func(int) int](Mod(2), Add(1)),
					strings.Title,
				),
				Flip2(strings.Repeat),
			)),
			slices.Collect[string],
			StringJoin("|"),
		)(in)

		exp := "Pipe|LikeLike|A|ProPro"

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		in := 33

		out := Pipe[func(int) int]()(in)

		if out != in {
			t.Errorf("wrong value %v, expected %v", out, in)
		}
	})
}
