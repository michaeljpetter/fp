package fp_test

import (
	"fmt"
	"math"
	"testing"

	. "github.com/michaeljpetter/fp"
)

func TestCond(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		clamp := Cond[int, int]{
			{Gt(100), Always[int](100)},
			{Lt(0), Always[int](0)},
			{T[int], Identity[int]},
		}.Apply

		cases := []struct {
			Name string
			In   int
			Exp  int
		}{
			{"Branch1", 231, 100},
			{"Branch2", -73, 0},
			{"BranchElse", 32, 32},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				out := clamp(c.In)

				if out != c.Exp {
					t.Errorf("wrong value %v, expected %v", out, c.Exp)
				}
			})
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		out := Cond[int, string]{
			{Gte(100), Always[int]("boiling")},
			{Lte(0), Always[int]("frozen")},
		}.Apply(44)

		exp := ""

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := Cond[string, float64](nil).Apply("irrelevant")

		exp := 0.

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})
}

func TestCond2(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		diff := Cond2[int, int, string]{
			{Gt2[int], func(a, b int) string { return fmt.Sprintf("bigger by %d", a-b) }},
			{Lt2[int], func(a, b int) string { return fmt.Sprintf("smaller by %d", b-a) }},
			{T2[int, int], Always2[int, int]("equal")},
		}.Apply

		cases := []struct {
			Name string
			In1  int
			In2  int
			Exp  string
		}{
			{"Branch1", 18, 7, "bigger by 11"},
			{"Branch2", -4, 13, "smaller by 17"},
			{"BranchElse", 32, 32, "equal"},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				out := diff(c.In1, c.In2)

				if out != c.Exp {
					t.Errorf("wrong value %v, expected %v", out, c.Exp)
				}
			})
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		out := Cond2[float64, int, string]{
			{math.IsInf, Always2[float64, int]("infinite!")},
		}.Apply(43, 1)

		exp := ""

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out := Cond2[string, int, float64](nil).Apply("irrelevant", 33)

		exp := 0.

		if out != exp {
			t.Errorf("wrong value %v, expected %v", out, exp)
		}
	})
}

func TestCondTo2(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		describe := CondTo2[int, string, int]{
			{Gt(0), func(a int) (string, int) { return "positive", a }},
			{Lt(0), func(a int) (string, int) { return "negative", -a }},
			{T[int], AlwaysTo2[int]("zero", 0)},
		}.Apply

		cases := []struct {
			Name string
			In   int
			Exp1 string
			Exp2 int
		}{
			{"Branch1", 42, "positive", 42},
			{"Branch2", -73, "negative", 73},
			{"BranchElse", -0, "zero", 0},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				out1, out2 := describe(c.In)

				if out1 != c.Exp1 || out2 != c.Exp2 {
					t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, c.Exp1, c.Exp2)
				}
			})
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		out1, out2 := CondTo2[int, string, float64]{
			{Gte(212), AlwaysTo2[int]("boiling", 212.)},
			{Lte(32), AlwaysTo2[int]("freezing", 32.)},
		}.Apply(44)

		exp1, exp2 := "", 0.

		if out1 != exp1 || out2 != exp2 {
			t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out1, out2 := CondTo2[string, float64, any](nil).Apply("irrelevant")

		exp1, exp2 := 0., any(nil)

		if out1 != exp1 || out2 != exp2 {
			t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
		}
	})
}

func TestCond2To2(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		diff := Cond2To2[int, int, string, int]{
			{Gt2[int], func(a, b int) (string, int) { return "bigger by", a - b }},
			{Lt2[int], func(a, b int) (string, int) { return "smaller by", b - a }},
			{T2[int, int], Always2To2[int, int]("equal", 0)},
		}.Apply

		cases := []struct {
			Name string
			In1  int
			In2  int
			Exp1 string
			Exp2 int
		}{
			{"Branch1", 18, 7, "bigger by", 11},
			{"Branch2", -4, 13, "smaller by", 17},
			{"BranchElse", 32, 32, "equal", 0},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				out1, out2 := diff(c.In1, c.In2)

				if out1 != c.Exp1 || out2 != c.Exp2 {
					t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, c.Exp1, c.Exp2)
				}
			})
		}
	})

	t.Run("NoMatch", func(t *testing.T) {
		out1, out2 := Cond2To2[float64, int, string, float64]{
			{math.IsInf, Always2To2[float64, int]("infinite!", 3.2)},
		}.Apply(43, 1)

		exp1, exp2 := "", 0.

		if out1 != exp1 || out2 != exp2 {
			t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
		}
	})

	t.Run("Empty", func(t *testing.T) {
		out1, out2 := Cond2To2[string, int, float64, any](nil).Apply("irrelevant", 33)

		exp1, exp2 := 0., any(nil)

		if out1 != exp1 || out2 != exp2 {
			t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
		}
	})
}
