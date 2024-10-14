// Package fp implements a set of operators and utilities
// used for functional programming in Go.
package fp

import (
	"errors"
	"fmt"
	"github.com/michaeljpetter/fp/types"
	"strings"
)

// Err is the root error wrapped by all errors in the fp package.
var Err = errors.New("fp")

// ErrArg indicates the use of an invalid argument.
var ErrArg = fmt.Errorf("%w: invalid argument", Err)

// ErrZero indicates the invalid use of a zero-valued argument.
var ErrZero = fmt.Errorf("%w: must be non-zero", ErrArg)

func newErrZero(name string) error {
	return fmt.Errorf("%w: %s", ErrZero, name)
}

// Identity simply returns its single input argument.
func Identity[T any](t T) T {
	return t
}

// StringLen wraps the [len] builtin for strings.
//
// [len]: https://pkg.go.dev/builtin#len
func StringLen(str string) int {
	return len(str)
}

// StringSplit wraps and curries [strings.Split] as data-last.
func StringSplit(sep string) func(string) []string {
	return func(str string) []string {
		return strings.Split(str, sep)
	}
}

// StringSplitN wraps and curries [strings.SplitN] as data-last.
func StringSplitN(sep string, n int) func(string) []string {
	return func(str string) []string {
		return strings.SplitN(str, sep, n)
	}
}

// StringJoin wraps and curries [strings.Join] as data-last.
func StringJoin(sep string) func([]string) string {
	return func(elems []string) string {
		return strings.Join(elems, sep)
	}
}

// SliceLen wraps the [len] builtin for slices.
//
// [len]: https://pkg.go.dev/builtin#len
func SliceLen[S ~[]E, E any](slice S) int {
	return len(slice)
}

// SliceAppend wraps the [append] builtin for slices.
// However, unlike the builtin, this method is fixed arity and only supports appending one element.
//
// [append]: https://pkg.go.dev/builtin#append
func SliceAppend[S ~[]E, E any](slice S, elem E) S {
	return append(slice, elem)
}

// Swap returns its two input arguments in swapped positions.
func Swap[T1, T2 any](t1 T1, t2 T2) (T2, T1) {
	return t2, t1
}

// Add wraps and curries the + operator.
func Add[T types.Ordered](t1 T) func(T) T {
	return func(t2 T) T { return t1 + t2 }
}

// Subtract wraps and curries the - operator, subtrahend first.
func Subtract[T types.Number](t2 T) func(T) T {
	return func(t1 T) T { return t1 - t2 }
}

// SubtractFrom wraps and curries the - operator, minuend first.
func SubtractFrom[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 - t2 }
}

// Multiply wraps and curries the * operator.
func Multiply[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 * t2 }
}

// Divide wraps and curries the / operator, dividend first.
func Divide[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 / t2 }
}

// DivideBy wraps and curries the / operator, divisor first.
func DivideBy[T types.Number](t2 T) func(T) T {
	return func(t1 T) T { return t1 / t2 }
}

// Mod wraps and curries the % operator, modulus first.
func Mod[T types.Int](t2 T) func(T) T {
	return func(t1 T) T { return t1 % t2 }
}

// ModOf wraps and curries the % operator, dividend first.
func ModOf[T types.Int](t1 T) func(T) T {
	return func(t2 T) T { return t1 % t2 }
}
