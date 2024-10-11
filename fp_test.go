package fp_test

import (
	"github.com/michaeljpetter/fp"
	"slices"
	"strings"
	"testing"
)

func TestIdentity(t *testing.T) {
	in := "the one"

	out := fp.Identity(in)

	if out != in {
		t.Errorf("wrong value %v, expected %v", out, in)
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
		t.Errorf("wrong values %v, expected %v", out, exp)
	}
}

func TestStringSplitN(t *testing.T) {
	in := "look-I'm-a-string"

	out := fp.StringSplitN("-", 3)(in)

	exp := strings.SplitN(in, "-", 3)

	if !slices.Equal(out, exp) {
		t.Errorf("wrong values %v, expected %v", out, exp)
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
		t.Errorf("wrong values %v, expected %v", out, exp)
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

func TestMultiply(t *testing.T) {
	out := fp.Multiply(4)(9)

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
