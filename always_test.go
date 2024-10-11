package fp_test

import (
	"github.com/michaeljpetter/fp"
	"testing"
)

func TestAlways(t *testing.T) {
	in := 77

	out := fp.Always(in)()

	if out != in {
		t.Errorf("wrong value %v, expected %v", out, in)
	}
}

func TestAlways2(t *testing.T) {
	in1, in2 := "pi", 3.14

	out1, out2 := fp.Always2(in1, in2)()

	if out1 != in1 || out2 != in2 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, in1, in2)
	}
}
