package fp_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestPipe(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		in := "pipe   like  a       pro"

		out := fp.Pipe[func(string) string](
			strings.Fields,
			slices.All[[]string],
			fp.Map2(fp.Pipe[func(int, string) string](
				fp.Project2(
					fp.Pipe[func(int) int](fp.Mod(2), fp.Add(1)),
					strings.Title,
				),
				fp.Flip2(strings.Repeat),
			)),
			slices.Collect[string],
			fp.StringJoin("|"),
		)(in)

		exp := "Pipe|LikeLike|A|ProPro"

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		in := 33

		out := fp.Pipe[func(int) int]()(in)

		if out != in {
			t.Errorf("wrong value %v, expected %v", out, in)
		}
	})
}
