package fp

import (
	"errors"
	"fmt"
	"github.com/michaeljpetter/fp/types"
	"strings"
)

var (
	Err     = errors.New("fp")
	ErrArg  = fmt.Errorf("%w: invalid argument", Err)
	ErrZero = fmt.Errorf("%w: must be non-zero", ErrArg)
)

func newErrZero(name string) error {
	return fmt.Errorf("%w: %s", ErrZero, name)
}

func Identity[T any](t T) T {
	return t
}

func StringLen(str string) int {
	return len(str)
}

func StringSplit(sep string) func(string) []string {
	return func(str string) []string {
		return strings.Split(str, sep)
	}
}

func StringSplitN(sep string, n int) func(string) []string {
	return func(str string) []string {
		return strings.SplitN(str, sep, n)
	}
}

func StringJoin(sep string) func([]string) string {
	return func(elems []string) string {
		return strings.Join(elems, sep)
	}
}

func SliceLen[S ~[]E, E any](slice S) int {
	return len(slice)
}

func SliceAppend[S ~[]E, E any](slice S, elem E) S {
	return append(slice, elem)
}

func Swap[T1, T2 any](t1 T1, t2 T2) (T2, T1) {
	return t2, t1
}

func Add[T types.Ordered](t1 T) func(T) T {
	return func(t2 T) T { return t1 + t2 }
}

func Subtract[T types.Number](t2 T) func(T) T {
	return func(t1 T) T { return t1 - t2 }
}

func SubtractFrom[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 - t2 }
}

func Multiply[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 * t2 }
}

func Divide[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 / t2 }
}

func DivideBy[T types.Number](t2 T) func(T) T {
	return func(t1 T) T { return t1 / t2 }
}

func Mod[T types.Int](t2 T) func(T) T {
	return func(t1 T) T { return t1 % t2 }
}

func ModOf[T types.Int](t1 T) func(T) T {
	return func(t2 T) T { return t1 % t2 }
}
