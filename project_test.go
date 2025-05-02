package fp_test

import (
	"math"
	"strings"
	"testing"

	"github.com/michaeljpetter/fp"
)

func TestProject2To2(t *testing.T) {
	in1, in2 := "e", 2.18

	out1, out2 := fp.Project2To2(strings.ToUpper, math.Floor)(in1, in2)

	exp1, exp2 := "E", 2.

	if out1 != exp1 || out2 != exp2 {
		t.Errorf("wrong elements (%v, %v), expected (%v, %v)", out1, out2, exp1, exp2)
	}
}
