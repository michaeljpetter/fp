package fp_test

import (
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestFlip2(t *testing.T) {
	in1, in2 := 3, "Hi"

	out := fp.Flip2(strings.Repeat)(in1, in2)

	exp := strings.Repeat(in2, in1)

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}

func TestFlip2To2(t *testing.T) {
	in1, in2 := "OK-", "OK-Cool"

	out1, out2 := fp.Flip2To2(strings.CutPrefix)(in1, in2)

	exp1, exp2 := strings.CutPrefix(in2, in1)

	if out1 != exp1 || out2 != exp2 {
		t.Errorf("wrong values (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
	}
}
