package fp_test

import (
	"github.com/michaeljpetter/fp"
	"strings"
	"testing"
)

func TestFlip2(t *testing.T) {
	in1, in2 := 3, "Hi"

	out := fp.Flip2(strings.Repeat)(in1, in2)

	exp := strings.Repeat(in2, in1)

	if out != exp {
		t.Errorf("wrong value %v, expected %v", out, exp)
	}
}
