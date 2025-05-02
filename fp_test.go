package fp_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestT(t *testing.T) {
	out := fp.T("irrelevant")

	exp := true

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestT2(t *testing.T) {
	out := fp.T2("irrelevant", 3.4)

	exp := true

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestF(t *testing.T) {
	out := fp.F("irrelevant")

	exp := false

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestF2(t *testing.T) {
	out := fp.F2("irrelevant", 3.4)

	exp := false

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestIdentity(t *testing.T) {
	in := "the one"

	out := fp.Identity(in)

	if out != in {
		t.Errorf("wrong value %v, expected %v", out, in)
	}
}

func TestIdentity2(t *testing.T) {
	in1, in2 := "the one", 22

	out1, out2 := fp.Identity2(in1, in2)

	if out1 != in1 || out2 != in2 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, in1, in2)
	}
}

func TestStringLen(t *testing.T) {
	in := "look I'm a string"

	out := fp.StringLen(in)

	if out != len(in) {
		t.Errorf("wrong len %v, expected %v", out, len(in))
	}
}

func TestStringSplit(t *testing.T) {
	in := "look I'm a string"

	out := fp.StringSplit(" ")(in)

	exp := strings.Split(in, " ")

	if !slices.Equal(out, exp) {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestStringSplitN(t *testing.T) {
	in := "look-I'm-a-string"

	out := fp.StringSplitN("-", 3)(in)

	exp := strings.SplitN(in, "-", 3)

	if !slices.Equal(out, exp) {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestStringJoin(t *testing.T) {
	in := []string{"look", "I'm", "a", "string"}

	out := fp.StringJoin("|")(in)

	exp := strings.Join(in, "|")

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestSliceLen(t *testing.T) {
	in := []int{1, 9, 3, 0}

	out := fp.SliceLen(in)

	if out != len(in) {
		t.Errorf("wrong len %v, expected %v", out, len(in))
	}
}

func TestSliceAppend(t *testing.T) {
	in := []int{1, 9}

	out := fp.SliceAppend(in, 7)

	exp := []int{1, 9, 7}

	if !slices.Equal(out, exp) {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestSwap(t *testing.T) {
	in1, in2 := "pi", 3.14

	out1, out2 := fp.Swap(in1, in2)

	if out1 != in2 || out2 != in1 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, in2, in1)
	}
}

func TestAdd(t *testing.T) {
	out := fp.Add(4)(-9)

	exp := -5

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestAdd2(t *testing.T) {
	out := fp.Add2(4, -9)

	exp := -5

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestSubtract(t *testing.T) {
	out := fp.Subtract(4)(9)

	exp := 5

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestSubtractFrom(t *testing.T) {
	out := fp.SubtractFrom(4)(9)

	exp := -5

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestSubtract2(t *testing.T) {
	out := fp.Subtract2(4, 9)

	exp := -5

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestMultiply(t *testing.T) {
	out := fp.Multiply(4)(9)

	exp := 36

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestMultiply2(t *testing.T) {
	out := fp.Multiply2(4, 9)

	exp := 36

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestDivide(t *testing.T) {
	out := fp.Divide(20)(5)

	exp := 4

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestDivideBy(t *testing.T) {
	out := fp.DivideBy(4)(20)

	exp := 5

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestDivide2(t *testing.T) {
	out := fp.Divide2(20, 5)

	exp := 4

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestMod(t *testing.T) {
	out := fp.Mod(4)(19)

	exp := 3

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestModOf(t *testing.T) {
	out := fp.ModOf(19)(5)

	exp := 4

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestMod2(t *testing.T) {
	out := fp.Mod2(19, 4)

	exp := 3

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestGt(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 2, 5, true},
		{"Less", 5, 2, false},
		{"Equal", 3, 3, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Gt(c.In1)(c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestGte(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 2, 5, true},
		{"Less", 5, 2, false},
		{"Equal", 3, 3, true},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Gte(c.In1)(c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestLt(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 2, 5, false},
		{"Less", 5, 2, true},
		{"Equal", 3, 3, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Lt(c.In1)(c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestLte(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 2, 5, false},
		{"Less", 5, 2, true},
		{"Equal", 3, 3, true},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Lte(c.In1)(c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestGt2(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 5, 2, true},
		{"Less", 2, 5, false},
		{"Equal", 3, 3, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Gt2(c.In1, c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestGte2(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 5, 2, true},
		{"Less", 2, 5, false},
		{"Equal", 3, 3, true},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Gte2(c.In1, c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestLt2(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 5, 2, false},
		{"Less", 2, 5, true},
		{"Equal", 3, 3, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Lt2(c.In1, c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}

func TestLte2(t *testing.T) {
	cases := []struct {
		Name string
		In1  int
		In2  int
		Exp  bool
	}{
		{"Greater", 5, 2, false},
		{"Less", 2, 5, true},
		{"Equal", 3, 3, true},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			out := fp.Lte2(c.In1, c.In2)

			if out != c.Exp {
				t.Errorf("wrong value %v, expected %v", out, c.Exp)
			}
		})
	}
}
