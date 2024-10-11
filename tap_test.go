package fp_test

import (
	"github.com/michaeljpetter/fp"
	"testing"
)

func TestTap(t *testing.T) {
	in := 77

	var arg int
	out := fp.Tap(func(a int) { arg = a })(in)

	if arg != in {
		t.Errorf("wrong arg %v, expected %v", arg, in)
	}
	if out != in {
		t.Errorf("wrong value %v, expected %v", out, in)
	}
}

func TestTap2(t *testing.T) {
	in1, in2 := "pi", 3.14

	var arg1 string
	var arg2 float64
	out1, out2 := fp.Tap2(func(a1 string, a2 float64) { arg1 = a1; arg2 = a2 })(in1, in2)

	if arg1 != in1 || arg2 != in2 {
		t.Errorf("wrong args (%v, %v), expected (%v, %v)", arg1, arg2, in1, in2)
	}
	if out1 != in1 || out2 != in2 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, in1, in2)
	}
}
