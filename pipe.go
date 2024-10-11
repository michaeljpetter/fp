package fp

import (
	"reflect"
	"slices"
)

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
