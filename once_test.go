package fp_test

import (
	"github.com/michaeljpetter/fp"
	"sync"
	"sync/atomic"
	"testing"
)

func TestOnce(t *testing.T) {
	in := func(i int) int { return i * i }

	out := fp.Once(in)

	if out(4) != 16 {
		t.Errorf("wrong value %v, expected %v", out(4), 16)
	}
	if out(7) != 16 {
		t.Errorf("wrong value %v, expected %v", out(7), 16)
	}
}

func TestOnceParallel(t *testing.T) {
	var calls atomic.Int32
	in := func() { calls.Add(1) }

	out := fp.Once(in)

	wg := new(sync.WaitGroup)
	wg.Add(8)

	for range 8 {
		go func() {
			defer wg.Done()
			out()
		}()
	}

	wg.Wait()

	if calls.Load() != 1 {
		t.Errorf("wrong number of calls %v, expected %v", calls.Load(), 1)
	}
}

func TestOncePanic(t *testing.T) {
	var calls atomic.Int32
	in := func() { calls.Add(1); panic("very bad!") }

	out := fp.Once(in)

	for range 3 {
		var p any
		func() {
			defer func() { p = recover() }()
			out()
		}()

		if p != "very bad!" {
			t.Errorf("wrong panic %v, expected %v", p, "very bad!")
		}
	}

	if calls.Load() != 1 {
		t.Errorf("wrong number of calls %v, expected %v", calls.Load(), 1)
	}
}
