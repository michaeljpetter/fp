// Package fp implements a set of operators and utilities
// used for functional programming in Go.
package fp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/michaeljpetter/fp/types"
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

// T is a unary predicate that always returns true.
func T[T any](T) bool {
	return true
}

// T2 is a binary predicate that always returns true.
func T2[T1, T2 any](T1, T2) bool {
	return true
}

// F is a unary predicate that always returns false.
func F[T any](T) bool {
	return false
}

// F2 is a binary predicate that always returns false.
func F2[T1, T2 any](T1, T2) bool {
	return false
}

// Identity simply returns its single input argument.
func Identity[T any](t T) T {
	return t
}

// Identity2 simply returns its two input arguments.
func Identity2[T1, T2 any](t1 T1, t2 T2) (T1, T2) {
	return t1, t2
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

// Add2 wraps the + operator.
func Add2[T types.Ordered](t1, t2 T) T {
	return t1 + t2
}

// Subtract wraps and curries the - operator, subtrahend first.
func Subtract[T types.Number](t2 T) func(T) T {
	return func(t1 T) T { return t1 - t2 }
}

// SubtractFrom wraps and curries the - operator, minuend first.
func SubtractFrom[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 - t2 }
}

// Subtract2 wraps the - operator.
func Subtract2[T types.Number](t1, t2 T) T {
	return t1 - t2
}

// Multiply wraps and curries the * operator.
func Multiply[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 * t2 }
}

// Multiply2 wraps the * operator.
func Multiply2[T types.Number](t1, t2 T) T {
	return t1 * t2
}

// Divide wraps and curries the / operator, dividend first.
func Divide[T types.Number](t1 T) func(T) T {
	return func(t2 T) T { return t1 / t2 }
}

// DivideBy wraps and curries the / operator, divisor first.
func DivideBy[T types.Number](t2 T) func(T) T {
	return func(t1 T) T { return t1 / t2 }
}

// Divide2 wraps the / operator.
func Divide2[T types.Number](t1, t2 T) T {
	return t1 / t2
}

// Mod wraps and curries the % operator, modulus first.
func Mod[T types.Int](t2 T) func(T) T {
	return func(t1 T) T { return t1 % t2 }
}

// ModOf wraps and curries the % operator, dividend first.
func ModOf[T types.Int](t1 T) func(T) T {
	return func(t2 T) T { return t1 % t2 }
}

// Mod2 wraps the % operator.
func Mod2[T types.Int](t1, t2 T) T {
	return t1 % t2
}

// Gt wraps and curries the > operator, in reverse operand order.
// e.g., Gt(3)(4) evaluates to 4 > 3.
func Gt[T types.Ordered](t2 T) func(T) bool {
	return func(t1 T) bool { return t1 > t2 }
}

// Gte wraps and curries the >= operator, in reverse operand order.
// e.g., Gte(3)(4) evaluates to 4 >= 3.
func Gte[T types.Ordered](t2 T) func(T) bool {
	return func(t1 T) bool { return t1 >= t2 }
}

// Lt wraps and curries the < operator, in reverse operand order.
// e.g., Lt(5)(2) evaluates to 2 < 5.
func Lt[T types.Ordered](t2 T) func(T) bool {
	return func(t1 T) bool { return t1 < t2 }
}

// Lte wraps and curries the <= operator, in reverse operand order.
// e.g., Lte(5)(2) evaluates to 2 <= 5.
func Lte[T types.Ordered](t2 T) func(T) bool {
	return func(t1 T) bool { return t1 <= t2 }
}

// Gt2 wraps the > operator.
// e.g., Gt2(4, 3) evaluates to 4 > 3.
func Gt2[T types.Ordered](t1, t2 T) bool {
	return t1 > t2
}

// Gte2 wraps the >= operator.
// e.g., Gte2(4, 3) evaluates to 4 >= 3.
func Gte2[T types.Ordered](t1, t2 T) bool {
	return t1 >= t2
}

// Lt2 wraps the < operator.
// e.g., Lt2(2, 5) evaluates to 2 < 5.
func Lt2[T types.Ordered](t1, t2 T) bool {
	return t1 < t2
}

// Lte2 wraps the <= operator.
// e.g., Lte2(2, 5) evaluates to 2 <= 5.
func Lte2[T types.Ordered](t1, t2 T) bool {
	return t1 <= t2
}
