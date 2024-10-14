package fp

import (
	"reflect"
	"slices"
)

// Pipe reduces a series of functions into a single function,
// passing the outputs of each function as the inputs to the next
// function, in first-to-last order.  The type parameter must be a
// function matching the inputs of the first given function and
// the outputs of the last given function.
func Pipe[F any](functions ...any) F {
	return reflect.MakeFunc(
		reflect.TypeFor[F](),
		func(args []reflect.Value) []reflect.Value {
			return Reduce(
				func(args []reflect.Value, function any) []reflect.Value {
					return reflect.ValueOf(function).Call(args)
				},
				args,
			)(slices.Values(functions))
		},
	).Interface().(F)
}
