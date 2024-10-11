package fp_test

import (
	"github.com/michaeljpetter/fp"
	"math"
	"strings"
	"testing"
)

func TestProject2(t *testing.T) {
	in1, in2 := "e", 2.18

	out1, out2 := fp.Project2(strings.ToUpper, math.Floor)(in1, in2)

	exp1, exp2 := "E", 2.

	if out1 != exp1 || out2 != exp2 {
		t.Errorf("wrong elements (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
	}
}
