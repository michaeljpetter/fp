package fp_test

import (
	"github.com/michaeljpetter/fp"
	"testing"
)

func TestNot(t *testing.T) {
	in := func(i int) bool { return 0 < i }

	out := fp.Not(in)

	if out(44) != !in(44) {
		t.Errorf("wrong value %v, expected %v", out(44), !in(44))
	}
	if out(-3) != !in(-3) {
		t.Errorf("wrong value %v, expected %v", out(-3), !in(-3))
	}
}

func TestNot2(t *testing.T) {
	in := func(i int, j int) bool { return i < j }

	out := fp.Not2(in)

	if out(3, 7) != !in(3, 7) {
		t.Errorf("wrong value %v, expected %v", out(3, 7), !in(3, 7))
	}
	if out(5, -5) != !in(5, -5) {
		t.Errorf("wrong value %v, expected %v", out(5, -5), !in(5, -5))
	}
}
