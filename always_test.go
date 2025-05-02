package fp_test

import (
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestAlways(t *testing.T) {
	in := 77

	out := fp.Always[string](in)("irrelevant")

	if out != in {
		t.Errorf("wrong value %v, expected %v", out, in)
	}
}

func TestAlways2(t *testing.T) {
	in := 77

	out := fp.Always2[string, float64](in)("irrelevant", 5.4)

	if out != in {
		t.Errorf("wrong value %v, expected %v", out, in)
	}
}

func TestAlwaysTo2(t *testing.T) {
	in1, in2 := 77, "ok"

	out1, out2 := fp.AlwaysTo2[string](in1, in2)("irrelevant")

	if out1 != in1 || out2 != in2 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, in1, in2)
	}
}

func TestAlways2To2(t *testing.T) {
	in1, in2 := 77, "ok"

	out1, out2 := fp.Always2To2[string, float64](in1, in2)("irrelevant", 5.4)

	if out1 != in1 || out2 != in2 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, in1, in2)
	}
}
